package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/vote"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/votetemplate"
)

type voteTemplateRepo struct {
	data *Data
}

var (
	ErrVoteTemplateNotFound = errors.New("votetemplate not found")
)

func NewVoteTemplateRepo(data *Data) biz.VoteTemplateRepo {
	return &voteTemplateRepo{
		data: data,
	}
}

func (repo voteTemplateRepo) FindAll(ctx context.Context) ([]*biz.VoteTemplateDto, error) {
	vt, err := repo.data.db.VoteTemplate.Query().
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	voteTemplates := make([]*biz.VoteTemplateDto, len(vt))
	for i, voteTemplate := range vt {
		voteTemplates[i] = new(biz.VoteTemplateDto).FromEntity(voteTemplate)
	}
	return voteTemplates, nil
}

func (repo voteTemplateRepo) Find(ctx context.Context, id int) (*biz.VoteTemplateDto, error) {
	vt, err := repo.data.db.VoteTemplate.
		Query().
		Where(votetemplate.ID(id)).
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.VoteTemplateDto).FromEntity(vt)
	return res, nil
}

func (repo voteTemplateRepo) FindByPoll(ctx context.Context, pollId int) ([]*biz.VoteTemplateDto, error) {
	vt, err := repo.data.db.VoteTemplate.
		Query().
		Where(votetemplate.HasPollWith(poll.ID(pollId))).
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	voteTemplates := make([]*biz.VoteTemplateDto, len(vt))
	for i, voteTemplate := range vt {
		voteTemplates[i] = new(biz.VoteTemplateDto).FromEntity(voteTemplate)
	}
	return voteTemplates, nil
}

func (repo voteTemplateRepo) Create(ctx context.Context, dto *biz.VoteTemplateCreateDto) (*biz.VoteTemplateDto, error) {
	tx, err := repo.data.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	rollback := func(tx *ent.Tx, err error) error {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	vt, err := tx.VoteTemplate.
		Create().
		SetAnswer(dto.Answer).
		SetPollID(dto.PollID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	createdVoteTemplate, err := repo.data.db.VoteTemplate.
		Query().
		Where(votetemplate.ID(vt.ID)).
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.VoteTemplateDto).FromEntity(createdVoteTemplate)
	return res, nil

}

func (repo voteTemplateRepo) Update(ctx context.Context, id int, dto *biz.VoteTemplateUpdateDto) (*biz.VoteTemplateDto, error) {
	vt, err := repo.data.db.VoteTemplate.UpdateOneID(id).SetAnswer(dto.Answer).Save(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.VoteTemplateDto).FromEntity(vt)
	return res, nil
}

func (repo voteTemplateRepo) Delete(ctx context.Context, id int) (*biz.VoteTemplateDto, error) {
	vt, err := repo.data.db.VoteTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = repo.data.db.VoteTemplate.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	// real-delete all votes beloning to the votetemplate
	_, err = repo.data.db.Vote.Delete().
		Where(vote.HasTemplateWith(votetemplate.ID(id))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.VoteTemplateDto).FromEntity(vt)
	return res, nil
}

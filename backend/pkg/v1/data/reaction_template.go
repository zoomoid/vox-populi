package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reaction"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reactiontemplate"
)

type reactionTemplateRepo struct {
	data *Data
}

var (
	ErrReactionTemplateNotFound = errors.New("reactiontemplate not found")
)

func NewReactionTemplateRepo(data *Data) biz.ReactionTemplateRepo {
	return &reactionTemplateRepo{
		data: data,
	}
}

func (repo reactionTemplateRepo) FindAll(ctx context.Context) ([]*biz.ReactionTemplateDto, error) {
	r, err := repo.data.db.ReactionTemplate.
		Query().
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	reactionTemplates := make([]*biz.ReactionTemplateDto, len(r))
	for i, reactionTemplate := range r {
		reactionTemplates[i] = new(biz.ReactionTemplateDto).FromEntity(reactionTemplate)
	}
	return reactionTemplates, nil
}

func (repo reactionTemplateRepo) Find(ctx context.Context, id int) (*biz.ReactionTemplateDto, error) {
	r, err := repo.data.db.ReactionTemplate.
		Query().
		Where(reactiontemplate.ID(id)).
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.ReactionTemplateDto).FromEntity(r)
	return res, nil
}

func (repo reactionTemplateRepo) FindByPoll(ctx context.Context, pollId int) ([]*biz.ReactionTemplateDto, error) {
	r, err := repo.data.db.ReactionTemplate.
		Query().
		Where(reactiontemplate.HasPollWith(poll.ID(pollId))).
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	reactionTemplates := make([]*biz.ReactionTemplateDto, len(r))
	for i, reactionTemplate := range r {
		reactionTemplates[i] = new(biz.ReactionTemplateDto).FromEntity(reactionTemplate)
	}
	return reactionTemplates, nil
}

func (repo reactionTemplateRepo) Create(ctx context.Context, dto *biz.ReactionTemplateCreateDto) (*biz.ReactionTemplateDto, error) {
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

	rt, err := tx.ReactionTemplate.
		Create().
		SetReaction(dto.Reaction).
		SetPollID(dto.PollID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	createdReactionTemplate, err := repo.data.db.ReactionTemplate.
		Query().
		Where(reactiontemplate.ID(rt.ID)).
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.ReactionTemplateDto).FromEntity(createdReactionTemplate)
	return res, nil

}

func (repo reactionTemplateRepo) Update(ctx context.Context, id int, dto *biz.ReactionTemplateUpdateDto) (*biz.ReactionTemplateDto, error) {
	rt, err := repo.data.db.ReactionTemplate.
		UpdateOneID(id).
		SetReaction(dto.Reaction).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.ReactionTemplateDto).FromEntity(rt)
	return res, nil
}

func (repo reactionTemplateRepo) Delete(ctx context.Context, id int) (*biz.ReactionTemplateDto, error) {
	rt, err := repo.data.db.ReactionTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = repo.data.db.ReactionTemplate.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	// cascade on all reactions
	_, err = repo.data.db.Reaction.Delete().
		Where(reaction.HasTemplateWith(reactiontemplate.ID(id))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.ReactionTemplateDto).FromEntity(rt)
	return res, nil
}

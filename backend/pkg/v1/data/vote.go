package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/vote"
)

var (
	ErrVoteNotFound = errors.New("vote not found")
)

type voteRepo struct {
	data *Data
}

func NewVoteRepo(data *Data) biz.VoteRepo {
	return &voteRepo{
		data: data,
	}
}

func (repo voteRepo) FindAll(ctx context.Context) ([]*biz.VoteDto, error) {
	v, err := repo.data.db.Vote.Query().
		Where(vote.IsDeleted(false)).
		WithTemplate().
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	votes := make([]*biz.VoteDto, len(v))
	for i, vote := range v {
		votes[i] = new(biz.VoteDto).FromEntity(vote)
	}
	return votes, nil
}

func (repo voteRepo) Find(ctx context.Context, id int) (*biz.VoteDto, error) {
	v, err := repo.data.db.Vote.
		Query().
		Where(vote.ID(id), vote.IsDeleted(false)).
		WithTemplate().
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.VoteDto).FromEntity(v)
	return res, nil
}

func (repo voteRepo) FindByPoll(ctx context.Context, pollId int) ([]*biz.VoteDto, error) {
	v, err := repo.data.db.Vote.
		Query().
		Where(vote.HasPollWith(poll.ID(pollId))).
		WithTemplate().
		All(ctx)
	if err != nil {
		return nil, err
	}

	votes := make([]*biz.VoteDto, len(v))
	for i, vote := range v {
		votes[i] = new(biz.VoteDto).FromEntity(vote)
	}
	return votes, nil
}

func (repo voteRepo) Create(ctx context.Context, dto *biz.VoteCreateDto) (*biz.VoteDto, error) {
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

	v, err := tx.Vote.
		Create().
		SetNonce(dto.Nonce).
		SetPollID(dto.PollID).
		SetTemplateID(dto.TemplateID).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	createdVote, err := repo.data.db.Vote.
		Query().
		Where(vote.ID(v.ID)).
		WithPoll().
		WithTemplate().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.VoteDto).FromEntity(createdVote)
	return res, nil
}

func (repo voteRepo) Update(ctx context.Context, id int, dto *biz.VoteUpdateDto) (*biz.VoteDto, error) {
	v, err := repo.data.db.Vote.
		Query().
		Where(vote.ID(id), vote.Nonce(dto.Nonce)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	_, err = v.Update().
		SetTemplateID(dto.TemplateID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.VoteDto).FromEntity(v)
	return res, nil
}

func (repo voteRepo) Delete(ctx context.Context, id int) (*biz.VoteDto, error) {
	v, err := repo.data.db.Vote.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if v.IsDeleted && v.DeletedAt != nil {
		return nil, ErrReactionNotFound
	}

	updatedVote, err := repo.data.db.Vote.
		UpdateOneID(id).
		SetIsDeleted(true).
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	res := new(biz.VoteDto).FromEntity(updatedVote)
	return res, nil
}

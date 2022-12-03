package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/reaction"
)

var (
	ErrReactionNotFound = errors.New("reaction not found")
)

type reactionRepo struct {
	data *Data
}

func NewReactionRepo(data *Data) biz.ReactionRepo {
	return &reactionRepo{
		data: data,
	}
}

func (repo reactionRepo) FindAll(ctx context.Context) ([]*biz.ReactionDto, error) {
	r, err := repo.data.db.Reaction.Query().
		WithTemplate().
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	reactions := make([]*biz.ReactionDto, len(r))
	for i, reaction := range r {
		reactions[i] = new(biz.ReactionDto).FromEntity(reaction)
	}
	return reactions, nil
}

func (repo reactionRepo) Find(ctx context.Context, id int) (*biz.ReactionDto, error) {
	r, err := repo.data.db.Reaction.
		Query().
		Where(reaction.ID(id)).
		WithTemplate().
		WithPoll().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.ReactionDto).FromEntity(r)
	return res, nil
}

func (repo reactionRepo) FindByPoll(ctx context.Context, pollId int) ([]*biz.ReactionDto, error) {
	r, err := repo.data.db.Reaction.
		Query().
		Where(reaction.HasPollWith(poll.ID(pollId))).
		WithTemplate().
		WithPoll().
		All(ctx)
	if err != nil {
		return nil, err
	}

	reactions := make([]*biz.ReactionDto, len(r))
	for i, reaction := range r {
		reactions[i] = new(biz.ReactionDto).FromEntity(reaction)
	}
	return reactions, nil
}

func (repo reactionRepo) Create(ctx context.Context, dto *biz.ReactionCreateDto) (*biz.ReactionDto, error) {
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

	r, err := tx.Reaction.
		Create().
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

	createdReaction, err := repo.data.db.Reaction.
		Query().
		Where(reaction.ID(r.ID)).
		WithPoll().
		WithTemplate().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	res := new(biz.ReactionDto).FromEntity(createdReaction)
	return res, nil
}

func (repo reactionRepo) Delete(ctx context.Context, id int) (*biz.ReactionDto, error) {
	r, err := repo.data.db.Reaction.
		Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = repo.data.db.Reaction.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.ReactionDto).FromEntity(r)
	return res, nil
}

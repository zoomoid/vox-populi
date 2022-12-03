package data

import (
	"context"
	"errors"
	"time"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
)

var (
	ErrCannotRepublishPoll = errors.New("cannot republish poll")
)

type pollRepo struct {
	data *Data
}

func NewPollRepo(data *Data) biz.PollRepo {
	return &pollRepo{
		data: data,
	}
}

func (repo pollRepo) FindAll(ctx context.Context) ([]*biz.PollDto, error) {
	p, err := repo.data.db.Poll.Query().
		Where(poll.IsDeleted(false)).
		WithVotes(func(vq *ent.VoteQuery) {
			vq.WithTemplate()
		}).
		WithReactions(func(rq *ent.ReactionQuery) {
			rq.WithTemplate()
		}).
		WithReactionTemplates().
		WithVoteTemplates().
		All(ctx)
	if err != nil {
		return nil, err
	}

	polls := make([]*biz.PollDto, len(p))
	for i, poll := range p {
		polls[i] = new(biz.PollDto).FromEntity(poll)
	}
	return polls, nil
}

func (repo pollRepo) FindById(ctx context.Context, id int) (*biz.PollDto, error) {
	p, err := repo.data.db.Poll.Query().
		Where(poll.ID(id), poll.IsDeleted(false)).
		WithVotes(func(vq *ent.VoteQuery) {
			vq.WithTemplate()
		}).
		WithReactions(func(rq *ent.ReactionQuery) {
			rq.WithTemplate()
		}).
		WithReactionTemplates().
		WithVoteTemplates().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	dto := new(biz.PollDto).FromEntity(p)
	return dto, nil
}

func (repo pollRepo) Find(ctx context.Context, preds ...predicate.Poll) ([]*biz.PollDto, error) {
	p, err := repo.data.db.Poll.Query().
		Where(poll.IsDeleted(false)). // TODO: re-evaluate if this is recommended if we want the most general query-ability here
		Where(preds...).
		WithVotes(func(vq *ent.VoteQuery) {
			vq.WithTemplate()
		}).
		WithReactions(func(rq *ent.ReactionQuery) {
			rq.WithTemplate()
		}).
		WithReactionTemplates().
		WithVoteTemplates().
		All(ctx)
	if err != nil {
		return nil, err
	}

	polls := make([]*biz.PollDto, len(p))
	for i, poll := range p {
		polls[i] = new(biz.PollDto).FromEntity(poll)
	}
	return polls, nil
}

func (repo pollRepo) Create(ctx context.Context, dto *biz.PollCreateDto) (*biz.PollDto, error) {
	tx, err := repo.data.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	poll, err := tx.Poll.
		Create().
		SetTitle(dto.Title).
		SetDescription(dto.Description).
		SetCanReact(dto.CanReact).
		SetCanVote(dto.CanVote).
		SetCanSeeResults(dto.CanSeeResults).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	for _, answer := range dto.Answers {
		_, err := tx.VoteTemplate.
			Create().
			SetAnswer(answer).
			SetPoll(poll).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, reaction := range dto.Reactions {
		_, err := tx.ReactionTemplate.
			Create().
			SetReaction(reaction).
			SetPoll(poll).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	// TODO(test): check if we need to propagate the back-reference
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	// TODO(test): check if we need to unwrap the poll object
	createdPoll, err := repo.data.db.Poll.Get(ctx, poll.ID)
	if err != nil {
		return nil, err
	}

	res := new(biz.PollDto).FromEntity(createdPoll)
	return res, nil
}

func (repo pollRepo) Update(ctx context.Context, id int, dto *biz.PollUpdateDto) (*biz.PollDto, error) {
	tx, err := repo.data.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	_, err = tx.Poll.UpdateOneID(id).
		SetTitle(dto.Title).
		SetDescription(dto.Description).
		SetCanReact(dto.CanReact).
		SetCanVote(dto.CanReact).
		SetCanSeeResults(dto.CanSeeResults).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	for _, vt := range dto.AddedAnswers {
		_, err := tx.VoteTemplate.
			Create().
			SetAnswer(vt).
			SetPollID(id).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, vt := range dto.UpdatedAnswers {
		_, err := tx.VoteTemplate.
			UpdateOneID(vt.ID).
			SetAnswer(vt.Answer).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, vt := range dto.DeletedAnswers {
		err := tx.VoteTemplate.DeleteOneID(vt).Exec(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, rt := range dto.AddedReactions {
		_, err := tx.ReactionTemplate.
			Create().
			SetReaction(rt).
			SetPollID(id).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, rt := range dto.UpdatedReactions {
		_, err := tx.ReactionTemplate.
			UpdateOneID(rt.ID).
			SetReaction(rt.Reaction).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	for _, rt := range dto.DeletedReactions {
		err := tx.ReactionTemplate.DeleteOneID(rt).Exec(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	updatedPoll, err := repo.data.db.Poll.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	res := new(biz.PollDto).FromEntity(updatedPoll)
	return res, nil
}

func (repo pollRepo) UpdateStatus(ctx context.Context, id int, dto *biz.PollStatusUpdateDto) (*biz.PollDto, error) {
	poll, err := repo.data.db.Poll.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if dto.IsLive && (poll.IsLive || poll.UnpublishedAt != nil) {
		oldPoll := new(biz.PollDto).FromEntity(poll)
		return oldPoll, ErrCannotRepublishPoll
	}

	pollUpdate := repo.data.db.Poll.UpdateOneID(id).SetIsLive(dto.IsLive)
	if dto.IsLive {
		pollUpdate.SetPublishedAt(time.Now())
	} else {
		pollUpdate.SetUnpublishedAt(time.Now())
	}
	poll, err = pollUpdate.Save(ctx)
	if err != nil {
		return nil, err
	}
	res := new(biz.PollDto).FromEntity(poll)
	return res, nil
}

func (repo pollRepo) Delete(ctx context.Context, id int) (*biz.PollDto, error) {
	poll, err := repo.data.db.Poll.
		UpdateOneID(id).
		SetIsDeleted(true).
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	res := new(biz.PollDto).FromEntity(poll)
	return res, nil
}

package services

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
)

type ReactionService struct {
	uc *biz.ReactionUsecase
}

func NewReactionService(r *biz.ReactionUsecase) *ReactionService {
	return &ReactionService{
		uc: r,
	}
}

func (r *ReactionService) List(ctx context.Context) ([]*biz.ReactionDto, error) {
	reactions, err := r.uc.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionService) Find(ctx context.Context, id int) (*biz.ReactionDto, error) {
	reaction, err := r.uc.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

func (r *ReactionService) FindByPoll(ctx context.Context, pollId int) ([]*biz.ReactionDto, error) {
	reactions, err := r.uc.FindByPoll(ctx, pollId)
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionService) Create(ctx context.Context, dto *biz.ReactionCreateDto) (*biz.ReactionDto, error) {
	reaction, err := r.uc.Create(ctx, dto)
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

func (r *ReactionService) Delete(ctx context.Context, id int) (*biz.ReactionDto, error) {
	reaction, err := r.uc.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

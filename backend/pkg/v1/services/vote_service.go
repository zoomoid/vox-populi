package services

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
)

type VoteService struct {
	uc *biz.VoteUsecase
}

func NewVoteService(r *biz.VoteUsecase) *VoteService {
	return &VoteService{
		uc: r,
	}
}

func (r *VoteService) Find(ctx context.Context, id int) (*biz.VoteDto, error) {
	vote, err := r.uc.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (r *VoteService) List(ctx context.Context) ([]*biz.VoteDto, error) {
	votes, err := r.uc.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (r *VoteService) FindByPoll(ctx context.Context, pollId int) ([]*biz.VoteDto, error) {
	votes, err := r.uc.FindByPoll(ctx, pollId)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (r *VoteService) Create(ctx context.Context, dto *biz.VoteCreateDto) (*biz.VoteDto, error) {
	vote, err := r.uc.Create(ctx, dto)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (r *VoteService) Update(ctx context.Context, id int, dto *biz.VoteUpdateDto) (*biz.VoteDto, error) {
	vote, err := r.uc.Update(ctx, id, dto)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (r *VoteService) Delete(ctx context.Context, id int) (*biz.VoteDto, error) {
	vote, err := r.uc.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

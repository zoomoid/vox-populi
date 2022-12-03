package services

import (
	"context"
	"errors"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/poll"
)

type PollService struct {
	puc *biz.PollUsecase
}

func NewPollService(puc *biz.PollUsecase) *PollService {
	return &PollService{
		puc: puc,
	}
}

func (p *PollService) Find(ctx context.Context, id int) (*biz.PollDto, error) {
	poll, err := p.puc.FindById(ctx, id)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return nil, err
	}
	return poll, nil
}

func (p *PollService) List(ctx context.Context) ([]*biz.PollDto, error) {
	polls, err := p.puc.FindAll(ctx)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return nil, err
	}
	return polls, nil
}

func (p *PollService) Create(ctx context.Context, dto *biz.PollCreateDto) (*biz.PollDto, error) {
	poll, err := p.puc.Create(ctx, dto)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return nil, err
	}
	return poll, nil
}

func (p *PollService) Current(ctx context.Context) (*biz.PollDto, error) {
	poll, err := p.puc.Find(ctx, poll.IsLive(true))
	if err != nil {
		return nil, err
	}
	if len(poll) > 1 {
		return nil, errors.New("live poll is ambiguous")
	}
	// There should always only be one
	return poll[0], nil
}

func (p *PollService) Update(ctx context.Context, id int, dto *biz.PollUpdateDto) (*biz.PollDto, error) {
	poll, err := p.puc.Update(ctx, id, dto)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return nil, err
	}
	return poll, nil
}

func (p *PollService) Delete(ctx context.Context, id int) (*biz.PollDto, error) {
	poll, err := p.puc.Delete(ctx, id)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return nil, err
	}
	return poll, nil
}

func (p *PollService) CanVote(ctx context.Context, id int) (bool, error) {
	poll, err := p.puc.FindById(ctx, id)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return false, err
	}
	return poll.CanVote, nil
}

func (p *PollService) CanReact(ctx context.Context, id int) (bool, error) {
	poll, err := p.puc.FindById(ctx, id)
	if err != nil {
		// TODO(refactor): replace this with a better user-facing error
		return false, err
	}
	return poll.CanReact, nil
}

func (p *PollService) Start(ctx context.Context, id int) error {
	_, err := p.puc.UpdateStatus(ctx, id, &biz.PollStatusUpdateDto{
		IsLive: true,
	})
	// TODO(refactor): replace this with a better user-facing error
	return err
}

func (p *PollService) Stop(ctx context.Context, id int) error {
	_, err := p.puc.UpdateStatus(ctx, id, &biz.PollStatusUpdateDto{
		IsLive: false,
	})
	// TODO(refactor): replace this with a better user-facing error
	return err
}

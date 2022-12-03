package biz

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
)

type ReactionTemplateDto struct {
	ID       int      `json:"id,omitempty"`
	Reaction string   `json:"reaction,omitempty"`
	Poll     *PollDto `json:"poll,omitempty"`
	MetadataMixin
}

func (rt ReactionTemplateDto) FromEntity(e *ent.ReactionTemplate) *ReactionTemplateDto {
	if e == nil {
		return nil
	}

	rt.ID = e.ID
	rt.Reaction = e.Reaction

	rt.CreatedAt = e.CreatedAt
	rt.UpdatedAt = e.UpdatedAt
	// rt.Revision = e.Revision

	rt.Poll = new(PollDto).FromEntity(e.Edges.Poll)

	return &rt
}

type ReactionTemplateCreateDto struct {
	Reaction string `json:"reaction,omitempty"`
	PollID   int    `json:"poll_id,omitempty"`
}

type ReactionTemplateUpdateDto struct {
	ID       int    `json:"id,omitempty"`
	Reaction string `json:"reaction,omitempty"`
}

type ReactionTemplateRepo interface {
	FindAll(context.Context) ([]*ReactionTemplateDto, error)
	Find(context.Context, int) (*ReactionTemplateDto, error)
	FindByPoll(context.Context, int) ([]*ReactionTemplateDto, error)

	Create(context.Context, *ReactionTemplateCreateDto) (*ReactionTemplateDto, error)
	Update(context.Context, int, *ReactionTemplateUpdateDto) (*ReactionTemplateDto, error)
	Delete(context.Context, int) (*ReactionTemplateDto, error)
}

type ReactionTemplateUsecase struct {
	repo ReactionTemplateRepo
}

func NewReactionTemplateUsecase(repo ReactionTemplateRepo) *ReactionTemplateUsecase {
	return &ReactionTemplateUsecase{repo: repo}
}

func (uc *ReactionTemplateUsecase) FindAll(ctx context.Context) ([]*ReactionTemplateDto, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *ReactionTemplateUsecase) Find(ctx context.Context, id int) (*ReactionTemplateDto, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *ReactionTemplateUsecase) FindByPoll(ctx context.Context, pollId int) ([]*ReactionTemplateDto, error) {
	return uc.repo.FindByPoll(ctx, pollId)
}

func (uc *ReactionTemplateUsecase) Create(ctx context.Context, dto *ReactionTemplateCreateDto) (*ReactionTemplateDto, error) {
	return uc.repo.Create(ctx, dto)
}

func (uc *ReactionTemplateUsecase) Update(ctx context.Context, id int, dto *ReactionTemplateUpdateDto) (*ReactionTemplateDto, error) {
	return uc.repo.Update(ctx, id, dto)
}

func (uc *ReactionTemplateUsecase) Delete(ctx context.Context, id int) (*ReactionTemplateDto, error) {
	return uc.repo.Delete(ctx, id)
}

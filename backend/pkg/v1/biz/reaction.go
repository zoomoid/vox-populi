package biz

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
)

type ReactionDto struct {
	ID       int                  `json:"id,omitempty"`
	Template *ReactionTemplateDto `json:"template,omitempty"`
	Poll     *PollDto             `json:"poll,omitempty"`
	MetadataMixin
}

func (r ReactionDto) FromEntity(e *ent.Reaction) *ReactionDto {
	if e == nil {
		return nil
	}

	r.ID = e.ID

	r.CreatedAt = e.CreatedAt
	r.UpdatedAt = e.UpdatedAt
	// r.Revision = e.Revision

	r.Template = new(ReactionTemplateDto).FromEntity(e.Edges.Template)
	r.Poll = new(PollDto).FromEntity(e.Edges.Poll)

	return &r
}

type ReactionCreateDto struct {
	TemplateID int `json:"template_id,omitempty"`
	PollID     int `json:"poll_id,omitempty"`
}

type ReactionRepo interface {
	FindAll(context.Context) ([]*ReactionDto, error)
	Find(context.Context, int) (*ReactionDto, error)
	FindByPoll(context.Context, int) ([]*ReactionDto, error)
	Create(context.Context, *ReactionCreateDto) (*ReactionDto, error)
	// Update(context.Context, string, *VoteUpdateDto) (*Vote, error)
	Delete(context.Context, int) (*ReactionDto, error)
}

type ReactionUsecase struct {
	repo ReactionRepo
}

func NewReactionUsecase(repo ReactionRepo) *ReactionUsecase {
	return &ReactionUsecase{repo: repo}
}

func (uc *ReactionUsecase) FindAll(ctx context.Context) ([]*ReactionDto, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *ReactionUsecase) Find(ctx context.Context, id int) (*ReactionDto, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *ReactionUsecase) FindByPoll(ctx context.Context, pollId int) ([]*ReactionDto, error) {
	return uc.repo.FindByPoll(ctx, pollId)
}

func (uc *ReactionUsecase) Create(ctx context.Context, dto *ReactionCreateDto) (*ReactionDto, error) {
	return uc.repo.Create(ctx, dto)
}

func (uc *ReactionUsecase) Delete(ctx context.Context, id int) (*ReactionDto, error) {
	return uc.repo.Delete(ctx, id)
}

package biz

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
)

type VoteTemplateDto struct {
	ID     int    `json:"id,omitempty"`
	Answer string `json:"answer,omitempty"`

	Poll *PollDto `json:"poll,omitempty"`
	MetadataMixin
}

func (vt VoteTemplateDto) FromEntity(e *ent.VoteTemplate) *VoteTemplateDto {
	if e == nil {
		return nil
	}

	vt.ID = e.ID
	vt.Answer = e.Answer

	vt.CreatedAt = e.CreatedAt
	vt.UpdatedAt = e.UpdatedAt
	// vt.Revision = e.Revision

	vt.Poll = new(PollDto).FromEntity(e.Edges.Poll)

	return &vt
}

type VoteTemplateCreateDto struct {
	Answer string `json:"answer,omitempty"`
	PollID int    `json:"poll_id,omitempty"`
}

type VoteTemplateUpdateDto struct {
	ID     int    `json:"id,omitempty"`
	Answer string `json:"answer,omitempty"`
}

type VoteTemplateRepo interface {
	FindAll(context.Context) ([]*VoteTemplateDto, error)
	Find(context.Context, int) (*VoteTemplateDto, error)
	FindByPoll(context.Context, int) ([]*VoteTemplateDto, error)

	Create(context.Context, *VoteTemplateCreateDto) (*VoteTemplateDto, error)
	Update(context.Context, int, *VoteTemplateUpdateDto) (*VoteTemplateDto, error)
	Delete(context.Context, int) (*VoteTemplateDto, error)
}

type VoteTemplateUsecase struct {
	repo VoteTemplateRepo
}

func NewVoteTemplateUsecase(repo VoteTemplateRepo) *VoteTemplateUsecase {
	return &VoteTemplateUsecase{repo: repo}
}

func (uc *VoteTemplateUsecase) FindAll(ctx context.Context) ([]*VoteTemplateDto, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *VoteTemplateUsecase) Find(ctx context.Context, id int) (*VoteTemplateDto, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *VoteTemplateUsecase) FindByPoll(ctx context.Context, pollId int) ([]*VoteTemplateDto, error) {
	return uc.repo.FindByPoll(ctx, pollId)
}

func (uc *VoteTemplateUsecase) Create(ctx context.Context, dto *VoteTemplateCreateDto) (*VoteTemplateDto, error) {
	return uc.repo.Create(ctx, dto)
}

func (uc *VoteTemplateUsecase) Update(ctx context.Context, id int, dto *VoteTemplateUpdateDto) (*VoteTemplateDto, error) {
	return uc.repo.Update(ctx, id, dto)
}

func (uc *VoteTemplateUsecase) Delete(ctx context.Context, id int) (*VoteTemplateDto, error) {
	return uc.repo.Delete(ctx, id)
}

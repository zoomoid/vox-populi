package biz

import (
	"context"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
)

type VoteDto struct {
	ID    int    `json:"id,omitempty"`
	Nonce string `json:"nonce,omitempty"`

	Template *VoteTemplateDto `json:"template,omitempty"`
	Poll     *PollDto         `json:"poll,omitempty"`
	MetadataMixin
}

func (v VoteDto) FromEntity(e *ent.Vote) *VoteDto {
	if e == nil {
		return nil
	}

	v.ID = e.ID
	v.Nonce = e.Nonce

	v.CreatedAt = e.CreatedAt
	v.UpdatedAt = e.UpdatedAt
	// v.Revision = e.Revision

	v.Poll = new(PollDto).FromEntity(e.Edges.Poll)
	v.Template = new(VoteTemplateDto).FromEntity(e.Edges.Template)

	return &v
}

type VoteCreateDto struct {
	Nonce      string `json:"nonce,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
	PollID     int    `json:"poll_id,omitempty"`
}

type VoteUpdateDto struct {
	Nonce      string `json:"nonce,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
	// PollID     int    `json:"poll_id,omitempty"`
}

type VoteRepo interface {
	FindAll(context.Context) ([]*VoteDto, error)
	Find(context.Context, int) (*VoteDto, error)
	FindByPoll(context.Context, int) ([]*VoteDto, error)
	Create(context.Context, *VoteCreateDto) (*VoteDto, error)
	Update(context.Context, int, *VoteUpdateDto) (*VoteDto, error)
	Delete(context.Context, int) (*VoteDto, error)
}

type VoteUsecase struct {
	repo VoteRepo
}

func NewVoteUsecase(repo VoteRepo) *VoteUsecase {
	return &VoteUsecase{repo: repo}
}

func (uc *VoteUsecase) FindAll(ctx context.Context) ([]*VoteDto, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *VoteUsecase) Find(ctx context.Context, id int) (*VoteDto, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *VoteUsecase) FindByPoll(ctx context.Context, pollId int) ([]*VoteDto, error) {
	return uc.repo.FindByPoll(ctx, pollId)
}

func (uc *VoteUsecase) Create(ctx context.Context, dto *VoteCreateDto) (*VoteDto, error) {
	return uc.repo.Create(ctx, dto)
}

func (uc *VoteUsecase) Update(ctx context.Context, id int, dto *VoteUpdateDto) (*VoteDto, error) {
	return uc.repo.Update(ctx, id, dto)
}

func (uc *VoteUsecase) Delete(ctx context.Context, id int) (*VoteDto, error) {
	return uc.repo.Delete(ctx, id)
}

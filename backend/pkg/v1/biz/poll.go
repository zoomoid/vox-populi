package biz

import (
	"context"
	"time"

	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data/ent/predicate"
)

type PollStatus struct {
	IsLive        bool       `json:"is_live,omitempty"`
	PublishedAt   *time.Time `json:"published_at,omitempty"`
	UnpublishedAt *time.Time `json:"unpublished_at,omitempty"`
}

type PollDto struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	CanVote       bool `json:"can_vote,omitempty"`
	CanReact      bool `json:"can_react,omitempty"`
	CanSeeResults bool `json:"can_see_results,omitempty"`

	MetadataMixin

	PollStatus

	VoteTemplates     []*VoteTemplateDto     `json:"vote_templates,omitempty"`
	Votes             []*VoteDto             `json:"votes,omitempty"`
	ReactionTemplates []*ReactionTemplateDto `json:"reaction_templates,omitempty"`
	Reactions         []*ReactionDto         `json:"reactions,omitempty"`
}

func (p PollDto) FromEntity(e *ent.Poll) *PollDto {
	if e == nil {
		return nil
	}

	votes := make([]*VoteDto, len(e.Edges.Votes))
	for i, ve := range e.Edges.Votes {
		votes[i] = new(VoteDto).FromEntity(ve)
	}

	voteTemplates := make([]*VoteTemplateDto, len(e.Edges.VoteTemplates))
	for i, vte := range e.Edges.VoteTemplates {
		voteTemplates[i] = new(VoteTemplateDto).FromEntity(vte)
	}

	reactions := make([]*ReactionDto, len(e.Edges.Reactions))
	for i, re := range e.Edges.Reactions {
		reactions[i] = new(ReactionDto).FromEntity(re)
	}

	reactionTemplates := make([]*ReactionTemplateDto, len(e.Edges.ReactionTemplates))
	for i, rte := range e.Edges.ReactionTemplates {
		reactionTemplates[i] = new(ReactionTemplateDto).FromEntity(rte)
	}

	p.ID = e.ID
	p.ID = e.ID
	p.Title = e.Title
	p.Description = e.Description
	p.IsLive = e.IsLive
	p.CanVote = e.CanVote
	p.CanReact = e.CanReact
	p.CanSeeResults = e.CanSeeResults

	p.CreatedAt = e.CreatedAt
	// p.Revision = e.Revision
	p.UpdatedAt = e.UpdatedAt

	p.PublishedAt = e.PublishedAt
	p.UnpublishedAt = e.UnpublishedAt

	p.Votes = votes
	p.VoteTemplates = voteTemplates
	p.Reactions = reactions
	p.ReactionTemplates = reactionTemplates

	return &p
}

type PollCreateDto struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	CanVote       bool `json:"can_vote,omitempty"`
	CanReact      bool `json:"can_react,omitempty"`
	CanSeeResults bool `json:"can_see_results,omitempty"`

	Answers   []string `json:"answers,omitempty"`
	Reactions []string `json:"reactions,omitempty"`
}

type PollUpdateDto struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	CanVote       bool `json:"can_vote,omitempty"`
	CanReact      bool `json:"can_react,omitempty"`
	CanSeeResults bool `json:"can_see_results,omitempty"`

	AddedAnswers   []string                 `json:"added_answers,omitempty"`
	UpdatedAnswers []*VoteTemplateUpdateDto `json:"updated_answers,omitempty"`
	DeletedAnswers []int                    `json:"deleted_answers,omitempty"`

	AddedReactions   []string                     `json:"added_reactions,omitempty"`
	UpdatedReactions []*ReactionTemplateUpdateDto `json:"updated_reactions,omitempty"`
	DeletedReactions []int                        `json:"deleted_reactions,omitempty"`
}

type PollStatusUpdateDto struct {
	IsLive bool `json:"is_live,omitempty"`
}

type PollRepo interface {
	FindAll(context.Context) ([]*PollDto, error)
	Find(context.Context, ...predicate.Poll) ([]*PollDto, error)
	FindById(context.Context, int) (*PollDto, error)
	Create(context.Context, *PollCreateDto) (*PollDto, error)
	Update(context.Context, int, *PollUpdateDto) (*PollDto, error)
	Delete(context.Context, int) (*PollDto, error)

	UpdateStatus(context.Context, int, *PollStatusUpdateDto) (*PollDto, error)
}

type PollUsecase struct {
	repo PollRepo
}

func NewPollUsecase(repo PollRepo) *PollUsecase {
	return &PollUsecase{repo: repo}
}

func (uc *PollUsecase) FindAll(ctx context.Context) ([]*PollDto, error) {
	return uc.repo.FindAll(ctx)
}

func (uc *PollUsecase) FindById(ctx context.Context, id int) (*PollDto, error) {
	return uc.repo.FindById(ctx, id)
}

func (uc *PollUsecase) Find(ctx context.Context, preds ...predicate.Poll) ([]*PollDto, error) {
	return uc.repo.Find(ctx, preds...)
}

func (uc *PollUsecase) Create(ctx context.Context, dto *PollCreateDto) (*PollDto, error) {
	return uc.repo.Create(ctx, dto)
}

func (uc *PollUsecase) Update(ctx context.Context, id int, dto *PollUpdateDto) (*PollDto, error) {
	return uc.repo.Update(ctx, id, dto)
}

func (uc *PollUsecase) UpdateStatus(ctx context.Context, id int, dto *PollStatusUpdateDto) (*PollDto, error) {
	return uc.repo.UpdateStatus(ctx, id, dto)
}

func (uc *PollUsecase) Delete(ctx context.Context, id int) (*PollDto, error) {
	return uc.repo.Delete(ctx, id)
}

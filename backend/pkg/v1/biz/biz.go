package biz

import "github.com/google/wire"

var Biz = wire.NewSet(
	NewPollUsecase,
	NewVoteUsecase,
	NewReactionUsecase,
	// NewVoteTemplateUsecase,
	// NewReactionTemplateUsecase,
)

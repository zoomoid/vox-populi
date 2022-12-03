package services

import (
	"github.com/google/wire"
)

var Services = wire.NewSet(
	NewPollService,
	NewVoteService,
	NewReactionService,
)

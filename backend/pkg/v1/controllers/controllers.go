package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var Controllers = wire.NewSet(
	NewPollController,
	NewVoteController,
	NewReactionController,
)

type Controller interface {
	Routes(*gin.RouterGroup) Controller
}

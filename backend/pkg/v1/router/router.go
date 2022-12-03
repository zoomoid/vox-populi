package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/controllers"
)

type Router struct {
	*gin.Engine
	poll     *controllers.PollController
	vote     *controllers.VoteController
	reaction *controllers.ReactionController
}

func NewRouter(
	pollCtrl *controllers.PollController,
	voteCtrl *controllers.VoteController,
	reactionCtrl *controllers.ReactionController,
) *Router {
	return &Router{
		poll:     pollCtrl,
		vote:     voteCtrl,
		reaction: reactionCtrl,
	}
}

func (r *Router) Routes(router *gin.RouterGroup) *Router {
	routes := router.Group("/v1")

	r.poll.Routes(routes)
	r.vote.Routes(routes)
	r.reaction.Routes(routes)

	return r
}

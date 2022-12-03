package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/services"
)

type ReactionController struct {
	service *services.ReactionService
}

var _ Controller = &ReactionController{}

func NewReactionController(svc *services.ReactionService) *ReactionController {
	ctrl := &ReactionController{
		service: svc,
	}

	return ctrl
}

func (ctrl *ReactionController) Routes(router *gin.RouterGroup) Controller {
	routes := router.Group("/reactions")

	routes.GET("/", ctrl.List)
	routes.GET("/by-poll/:poll-id", ctrl.GetByPoll)
	routes.GET("/:reaction-id", ctrl.Get)
	routes.POST("/", ctrl.Create)
	routes.DELETE("/:reaction-id", ctrl.Delete)

	return ctrl
}

func (ctrl ReactionController) List(c *gin.Context) {
	reactions, err := ctrl.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reactions)
	}
}

func (ctrl ReactionController) Get(c *gin.Context) {
	is := c.Param("reaction-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reaction, err := ctrl.service.Find(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reaction)
	}
}

func (ctrl ReactionController) GetByPoll(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reactions, err := ctrl.service.FindByPoll(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reactions)
	}
}

func (ctrl ReactionController) Create(c *gin.Context) {
	reactionCreateDto := &biz.ReactionCreateDto{}
	if err := c.ShouldBindJSON(reactionCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reaction, err := ctrl.service.Create(c.Request.Context(), reactionCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reaction)
	}
}

func (ctrl ReactionController) Delete(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reaction, err := ctrl.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reaction)
	}
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/services"
)

type VoteController struct {
	service *services.VoteService
}

var _ Controller = &VoteController{}

func NewVoteController(svc *services.VoteService) *VoteController {
	ctrl := &VoteController{
		service: svc,
	}

	return ctrl
}

func (ctrl *VoteController) Routes(router *gin.RouterGroup) Controller {
	routes := router.Group("/votes")

	routes.GET("/", ctrl.List)
	routes.GET("/by-poll/:poll-id", ctrl.GetByPoll)
	routes.GET("/:vote-id", ctrl.Get)
	routes.POST("/", ctrl.Create)
	routes.PATCH("/:vote-id", ctrl.Update)
	routes.DELETE("/:vote-id", ctrl.Delete)

	return ctrl
}

func (ctrl VoteController) List(c *gin.Context) {
	reactions, err := ctrl.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, reactions)
	}
}

func (ctrl VoteController) Get(c *gin.Context) {
	is := c.Param("vote-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vote, err := ctrl.service.Find(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

func (ctrl VoteController) GetByPoll(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	votes, err := ctrl.service.FindByPoll(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, votes)
	}
}

func (ctrl VoteController) Create(c *gin.Context) {
	voteCreateDto := &biz.VoteCreateDto{}
	if err := c.ShouldBindJSON(voteCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vote, err := ctrl.service.Create(c.Request.Context(), voteCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

func (ctrl VoteController) Update(c *gin.Context) {
	is := c.Param("vote-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	voteUpdateDto := &biz.VoteUpdateDto{}
	if err := c.ShouldBindJSON(voteUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vote, err := ctrl.service.Update(c.Request.Context(), id, voteUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, vote)
	}
}

func (ctrl VoteController) Delete(c *gin.Context) {
	is := c.Param("vote-id")
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

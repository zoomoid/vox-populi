package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/services"
)

type PollController struct {
	service *services.PollService
}

var _ Controller = &PollController{}

func NewPollController(svc *services.PollService) *PollController {
	ctrl := &PollController{
		service: svc,
	}

	return ctrl
}

func (ctrl *PollController) Routes(router *gin.RouterGroup) Controller {
	routes := router.Group("/polls")

	routes.GET("/", ctrl.List)
	routes.GET("/current", ctrl.GetCurrent)
	routes.GET("/:poll-id", ctrl.Get)
	routes.POST("/", ctrl.Create)
	routes.PATCH("/:poll-id", ctrl.Update)
	routes.DELETE("/:poll-id", ctrl.Delete)
	routes.POST("/:poll-id/start", ctrl.Start)
	routes.POST("/:poll-id/stop", ctrl.Stop)

	return ctrl
}

func (ctrl PollController) List(c *gin.Context) {
	polls, err := ctrl.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, polls)
	}
}

func (ctrl PollController) Get(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	poll, err := ctrl.service.Find(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, poll)
	}
}

func (ctrl PollController) Create(c *gin.Context) {
	pollCreateDto := &biz.PollCreateDto{}
	if err := c.ShouldBindJSON(pollCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poll, err := ctrl.service.Create(c.Request.Context(), pollCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, poll)
	}
}

func (ctrl PollController) Update(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pollUpdateDto := &biz.PollUpdateDto{}
	if err := c.ShouldBindJSON(pollUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	poll, err := ctrl.service.Update(c.Request.Context(), id, pollUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, poll)
	}
}

func (ctrl PollController) Delete(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	poll, err := ctrl.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, poll)
	}
}

func (ctrl PollController) GetCurrent(c *gin.Context) {
	poll, err := ctrl.service.Current(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, poll)
	}
}

func (ctrl PollController) Start(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.service.Start(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{})
	}
}

func (ctrl PollController) Stop(c *gin.Context) {
	is := c.Param("poll-id")
	id, err := strconv.Atoi(is)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.service.Stop(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{})
	}
}

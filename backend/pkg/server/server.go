package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apiv1router "github.com/zoomoid/vox-populi/backend/pkg/v1/router"
)

type ApiRouter struct {
	Engine *gin.Engine
	Routes *gin.RouterGroup
}

func NewServer(v1 *apiv1router.Router) *ApiRouter {
	e := gin.Default()

	e.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hahn oder ball?")
	})

	// TODO(feat): add healthcheck handlers to root router

	r := e.Group("/api")
	v1.Routes(r)

	return &ApiRouter{e, r}
}

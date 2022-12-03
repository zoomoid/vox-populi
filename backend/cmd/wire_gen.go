// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/zoomoid/vox-populi/backend/cmd/options"
	"github.com/zoomoid/vox-populi/backend/pkg/server"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/controllers"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/router"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/services"
)

// Injectors from wire.go:

func wireApp(configuration *options.Configuration) (*server.ApiRouter, func(), error) {
	dataData, cleanup, err := data.NewData(configuration)
	if err != nil {
		return nil, nil, err
	}
	pollRepo := data.NewPollRepo(dataData)
	pollUsecase := biz.NewPollUsecase(pollRepo)
	pollService := services.NewPollService(pollUsecase)
	pollController := controllers.NewPollController(pollService)
	voteRepo := data.NewVoteRepo(dataData)
	voteUsecase := biz.NewVoteUsecase(voteRepo)
	voteService := services.NewVoteService(voteUsecase)
	voteController := controllers.NewVoteController(voteService)
	reactionRepo := data.NewReactionRepo(dataData)
	reactionUsecase := biz.NewReactionUsecase(reactionRepo)
	reactionService := services.NewReactionService(reactionUsecase)
	reactionController := controllers.NewReactionController(reactionService)
	routerRouter := router.NewRouter(pollController, voteController, reactionController)
	apiRouter := server.NewServer(routerRouter)
	return apiRouter, func() {
		cleanup()
	}, nil
}

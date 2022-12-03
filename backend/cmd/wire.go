//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"

	"github.com/zoomoid/vox-populi/backend/pkg"
	"github.com/zoomoid/vox-populi/backend/pkg/config"
	"github.com/zoomoid/vox-populi/backend/pkg/server"
)

func wireApp(*config.Configuration) (*server.ApiRouter, func(), error) {
	panic(wire.Build(
		pkg.Modules,
	))
}

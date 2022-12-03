package pkg

import (
	"github.com/google/wire"
	"github.com/zoomoid/vox-populi/backend/pkg/server"
	v1 "github.com/zoomoid/vox-populi/backend/pkg/v1"
)

var Modules = wire.NewSet(
	server.NewServer,
	v1.Modules,
)

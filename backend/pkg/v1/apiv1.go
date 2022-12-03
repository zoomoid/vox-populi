package v1

import (
	"github.com/google/wire"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/biz"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/controllers"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/data"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/router"
	"github.com/zoomoid/vox-populi/backend/pkg/v1/services"
)

var Modules = wire.NewSet(
	controllers.Controllers,
	services.Services,
	data.Repositories,
	biz.Biz,
	router.NewRouter,
)

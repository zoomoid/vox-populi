package logging

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/zoomoid/qotd/pkg/config"
	"github.com/zoomoid/qotd/pkg/di/errors"
)

func New(cfg *config.Config) (*logr.Logger, error) {
	if cfg == nil {
		return nil, errors.ErrConfigIsNil
	}

	level := zerolog.Level(cfg.Logging.Level)

	zerolog.SetGlobalLevel(level)
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"

	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()

	log := zerologr.New(&zl)

	return &log, nil
}

package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/zoomoid/qotd/pkg/config"
	"github.com/zoomoid/qotd/pkg/logging"
	"github.com/zoomoid/qotd/pkg/metrics"
	"go.uber.org/dig"

	_ "github.com/zoomoid/qotd/pkg/flags"

	flag "github.com/spf13/pflag"
)

func main() {

	flag.Parse()

	zl := zerolog.New(os.Stderr)
	zl = zl.With().Caller().Timestamp().Logger()

	setupLogger := zerologr.New(&zl)

	c := dig.New()
	c.Provide(config.New(flag.CommandLine))

	err := c.Provide(logging.New)
	if err != nil {
		setupLogger.Error(err, "failed to provide logging module")
		panic(err)
	}

	err = c.Provide(metrics.New)
	if err != nil {
		setupLogger.Error(err, "failed to provide metrics module")
		panic(err)
	}

	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

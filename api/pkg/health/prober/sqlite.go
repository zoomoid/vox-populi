package prober

import (
	"context"
	"errors"

	"github.com/zoomoid/qotd/pkg/connectors/sqlite"
	"github.com/zoomoid/qotd/pkg/metrics"
)

type SQLiteProbe struct {
	config  *probeConfig
	client  *sqlite.Client
	metrics *metrics.ProbeMetrics
	context context.Context
}

var _ prober = &SQLiteProbe{}

func NewSQLiteProbe(ctx context.Context, conn *sqlite.Client, cfg *probeConfig, metrics *metrics.ProbeMetrics) (*SQLiteProbe, error) {

	metrics.FailureThresholdTotal.Set(float64(cfg.failureThreshold))
	metrics.SuccessThresholdTotal.Set(float64(cfg.failureThreshold))
	metrics.InitialDelayDurationSeconds.Set(float64(cfg.initialDelay))
	metrics.TimeoutDurationSeconds.Set(float64(cfg.timeout))
	metrics.PeriodDurationSeconds.Set(float64(cfg.period))

	p := &SQLiteProbe{
		client:  conn,
		config:  cfg,
		metrics: metrics,
		context: ctx,
	}

	return p, nil
}

func (p *SQLiteProbe) Probe() (Result, error) {
	probeContext, cancel := context.WithTimeout(p.context, p.config.timeout)
	defer cancel()
	err := p.client.DB().PingContext(probeContext)

	if errors.Is(err, context.DeadlineExceeded) {
		p.metrics.TimeoutsTotal.Inc()
		return Failure, err
	} else if err != nil {
		p.metrics.FailuresTotal.Inc()
		return Failure, err
	} else {
		p.metrics.SuccessTotal.Inc()
		return Success, nil
	}
}

func (p *SQLiteProbe) Name() string {
	return "sqlite"
}

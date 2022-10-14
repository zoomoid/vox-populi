package prober

import (
	"context"
	"errors"

	"github.com/zoomoid/qotd/pkg/connectors/redis"
	"github.com/zoomoid/qotd/pkg/metrics"
)

type RedisProbe struct {
	config  *probeConfig
	client  *redis.Client
	metrics *metrics.ProbeMetrics
	context context.Context
}

func NewRedisProbe(ctx context.Context, conn *redis.Client, cfg *probeConfig, metrics *metrics.ProbeMetrics) (*RedisProbe, error) {
	metrics.FailureThresholdTotal.Set(float64(cfg.failureThreshold))
	metrics.SuccessThresholdTotal.Set(float64(cfg.failureThreshold))
	metrics.InitialDelayDurationSeconds.Set(float64(cfg.initialDelay))
	metrics.TimeoutDurationSeconds.Set(float64(cfg.timeout))
	metrics.PeriodDurationSeconds.Set(float64(cfg.period))

	p := &RedisProbe{
		client:  conn,
		config:  cfg,
		metrics: metrics,
		context: ctx,
	}

	return p, nil
}

var _ prober = &RedisProbe{}

func (p *RedisProbe) Probe() (Result, error) {
	probeContext, cancel := context.WithTimeout(p.context, p.config.timeout)
	defer cancel()

	err := p.client.DB().Ping(probeContext).Err()

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

func (p *RedisProbe) Name() string {
	return "memcached"
}

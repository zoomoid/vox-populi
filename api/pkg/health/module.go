package health

import (
	"context"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/zoomoid/qotd/pkg/config"
	"github.com/zoomoid/qotd/pkg/connectors/redis"
	"github.com/zoomoid/qotd/pkg/connectors/sqlite"
	"github.com/zoomoid/qotd/pkg/metrics"

	"github.com/zoomoid/qotd/pkg/di/errors"
)

type HealthModule struct {
	sqliteConn *sqlite.Client
	redisConn  *redis.Client
	cfg        *config.Config
	metrics    *metrics.MetricsModule
	logger     *logr.Logger

	probeManager *probeManager
}

func NewHealthModule(logger *logr.Logger, sqliteConn *sqlite.Client, redisConn *redis.Client, cfg *config.Config, metrics *metrics.MetricsModule) (*HealthModule, error) {
	if sqliteConn == nil {
		return nil, errors.ErrSqliteClientIsNil
	}
	if redisConn == nil {
		return nil, errors.ErrRedisClientIsNil
	}
	if cfg == nil {
		return nil, errors.ErrConfigIsNil
	}
	if metrics == nil {
		return nil, errors.ErrMetricsIsNil
	}

	l := logger.WithName("health_module")

	probeMgr := &probeManager{
		workers:        make(map[string]*worker),
		start:          time.Now(),
		resultsManager: NewResultsManager(),
		logger:         &l,
	}

	module := &HealthModule{
		sqliteConn: sqliteConn,
		redisConn:  redisConn,
		cfg:        cfg,
		metrics:    metrics,

		logger:       &l,
		probeManager: probeMgr,
	}

	sqliteProbe, sqliteCfg, err := module.sqlite()
	if err != nil {
		return nil, err
	}
	l.V(3).Info("Created SQLite probe")

	redisProbe, redisCfg, err := module.redis()
	if err != nil {
		return nil, err
	}
	l.V(3).Info("Created Redis probe")

	probeMgr.Add(sqliteProbe, sqliteCfg)
	l.V(3).Info("Installed SQLite probe")

	probeMgr.Add(redisProbe, redisCfg)
	l.V(3).Info("Installed Redis probe")

	return module, nil
}

func (m *HealthModule) sqlite() (*SQLiteProbe, *probeConfig, error) {

	cfg := m.cfg.Probes

	sqliteCfg := &probeConfig{
		initialDelay:     time.Duration(cfg.SQLite.InitialDelaySeconds) * time.Second,
		period:           time.Duration(cfg.SQLite.PeriodSeconds) * time.Second,
		timeout:          time.Duration(cfg.SQLite.TimeoutSeconds) * time.Second,
		successThreshold: cfg.SQLite.SuccessThreshold,
		failureThreshold: cfg.SQLite.FailureThreshold,
	}

	sqliteProbeMetrics := m.metrics.Get().Probes("sqlite")
	sqliteProbe, err := NewSQLiteProbe(context.Background(), m.sqliteConn, sqliteCfg, sqliteProbeMetrics)
	if err != nil {
		return nil, nil, err
	}

	return sqliteProbe, sqliteCfg, nil
}

func (m *HealthModule) redis() (*RedisProbe, *probeConfig, error) {

	cfg := m.cfg.Probes

	redisCfg := &probeConfig{
		initialDelay:     time.Duration(cfg.SQLite.InitialDelaySeconds) * time.Second,
		period:           time.Duration(cfg.SQLite.PeriodSeconds) * time.Second,
		timeout:          time.Duration(cfg.SQLite.TimeoutSeconds) * time.Second,
		successThreshold: cfg.SQLite.SuccessThreshold,
		failureThreshold: cfg.SQLite.FailureThreshold,
	}

	redisProbeMetrics := m.metrics.Get().Probes("redis")
	redisProbe, err := NewRedisProbe(context.Background(), m.redisConn, redisCfg, redisProbeMetrics)
	if err != nil {
		return nil, nil, err
	}

	return redisProbe, redisCfg, nil
}

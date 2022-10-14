package redis

import (
	"fmt"

	"github.com/go-logr/logr"
	redis "github.com/go-redis/redis/v9"
	"github.com/zoomoid/qotd/pkg/config"
	dierrors "github.com/zoomoid/qotd/pkg/di/errors"
)

type Client struct {
	db     *redis.Client
	logger *logr.Logger
	config *config.Config
}

var ()

func New(l *logr.Logger, cfg *config.Config) (*Client, error) {
	if l == nil {
		return nil, dierrors.ErrLoggerIsNil
	}
	if cfg == nil {
		return nil, dierrors.ErrConfigIsNil
	}

	addr := fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port)
	opts := &redis.Options{
		Addr:     addr,
		DB:       cfg.Cache.DB,
		Username: cfg.Cache.Username,
		Password: cfg.Cache.Password,
	}
	if cfg.Cache.URL != "" {
		var err error
		opts, err = redis.ParseURL(cfg.Cache.URL)
		if err != nil {
			return nil, err
		}
		l.V(3).Info("Parsed Redis URL", "url", cfg.Cache.URL)
	}

	rdb := redis.NewClient(opts)

	return &Client{
		logger: l,
		config: cfg,
		db:     rdb,
	}, nil
}

func (c *Client) DB() *redis.Client {
	return c.db
}

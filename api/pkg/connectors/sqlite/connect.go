package sqlite

import (
	"database/sql"
	"errors"

	"github.com/go-logr/logr"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zoomoid/qotd/pkg/config"
	dierrors "github.com/zoomoid/qotd/pkg/di/errors"
)

type Client struct {
	db     *sql.DB
	logger *logr.Logger
	config *config.Config
}

func New(l *logr.Logger, cfg *config.Config) (*Client, error) {
	if l == nil {
		return nil, dierrors.ErrLoggerIsNil
	}
	if cfg == nil {
		return nil, dierrors.ErrConfigIsNil
	}

	path := cfg.Database.SQLite.Path

	if path == "" {
		return nil, errors.New("path is required")
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	l.V(3).Info("Connected to SQLite database", "path", path)

	return &Client{
		logger: l,
		config: cfg,
		db:     db,
	}, nil
}

// TODO(zoomoid): this is not nil-safe
func (c *Client) DB() *sql.DB {
	return c.db
}

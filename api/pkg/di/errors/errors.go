package errors

import "errors"

var (
	ErrConfigIsNil error = errors.New("config dependency is nil")

	ErrLoggerIsNil error = errors.New("logger dependency is nil")

	ErrRedisClientIsNil  error = errors.New("redis client dependency is nil")
	ErrSqliteClientIsNil error = errors.New("sqlite client dependency is nil")

	ErrMetricsIsNil error = errors.New("metrics dependency is nil")
)

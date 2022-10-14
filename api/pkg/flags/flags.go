package flags

import (
	flag "github.com/spf13/pflag"
	"github.com/zoomoid/qotd/pkg/config/defaults"
	n "github.com/zoomoid/qotd/pkg/flags/names"
)

var (
	host string
	port uint16

	verbosity int

	sqliteTimeout          int64
	sqlitePeriod           int64
	sqliteInitialDelay     int64
	sqliteFailureThreshold int64
	sqliteSuccessThreshold int64

	memcachedTimeout          int64
	memcachedPeriod           int64
	memcachedInitialDelay     int64
	memcachedFailureThreshold int64
	memcachedSuccessThreshold int64
)

func init() {
	flag.StringVar(&host, n.Host, defaults.Host, n.HostDescription)
	flag.Uint16Var(&port, n.Port, defaults.Port, n.PortDescrition)

	flag.IntVarP(&verbosity, n.Verbosity, n.VerbosityShort, defaults.Verbosity, n.VerbosityDescription)

	flag.Int64Var(&sqliteTimeout, n.SqliteTimeout, defaults.SqliteTimeout, n.SqliteTimeoutDescription)
	flag.Int64Var(&sqlitePeriod, n.SqlitePeriod, defaults.SqlitePeriod, n.SqlitePeriodDescription)
	flag.Int64Var(&sqliteInitialDelay, n.SqliteInitialDelay, defaults.SqliteInitialDelay, n.SqliteInitialDelayDescription)
	flag.Int64Var(&sqliteFailureThreshold, n.SqliteFailureThreshold, defaults.SqliteFailureThreshold, n.SqliteFailureThresholdDescription)
	flag.Int64Var(&sqliteSuccessThreshold, n.SqliteSuccessThreshold, defaults.SqliteSuccessThreshold, n.SqliteSuccessThresholdDescription)

	flag.Int64Var(&memcachedTimeout, n.RedisTimeout, defaults.RedisTimeout, n.RedisTimeoutDescription)
	flag.Int64Var(&memcachedPeriod, n.RedisPeriod, defaults.RedisPeriod, n.RedisPeriodDescription)
	flag.Int64Var(&memcachedInitialDelay, n.RedisInitialDelay, defaults.RedisInitialDelay, n.RedisInitialDelayDescription)
	flag.Int64Var(&memcachedFailureThreshold, n.RedisFailureThreshold, defaults.RedisFailureThreshold, n.RedisFailureThresholdDescription)
	flag.Int64Var(&memcachedSuccessThreshold, n.RedisSuccessThreshold, defaults.RedisSuccessThreshold, n.RedisSuccessThresholdDescription)
}

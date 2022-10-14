package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zoomoid/qotd/pkg/config/defaults"
)

func New(flags *pflag.FlagSet) func() *Config {
	return func() *Config {
		viper.SetEnvPrefix("QOTD_")

		setDefaults()
		bindEnv()
		bindFlags(flags)

		cfg := &Cfg{
			Logging: Logging{
				Level: viper.GetInt("logging.level"),
			},
			Server: Server{
				Host: viper.GetString("server.host"),
				Port: viper.GetUint16("server.port"),
			},
			Database: Database{
				Backend: DatabaseBackend(viper.GetString("database.backend")),
				SQLite: SQLite{
					Path: viper.GetString("database.sqlite.path"),
				},
			},
			Probes: Probes{
				SQLite: ProbeConfiguration{
					TimeoutSeconds:      viper.GetInt64("probes.sqlite.timeoutSeconds"),
					PeriodSeconds:       viper.GetInt64("probes.sqlite.periodSeconds"),
					InitialDelaySeconds: viper.GetInt64("probes.sqlite.initialDelaySeconds"),
					FailureThreshold:    viper.GetInt64("probes.sqlite.failureThreshold"),
					SuccessThreshold:    viper.GetInt64("probes.sqlite.successThreshold"),
				},
				Redis: ProbeConfiguration{
					TimeoutSeconds:      viper.GetInt64("probes.redis.timeoutSeconds"),
					PeriodSeconds:       viper.GetInt64("probes.redis.periodSeconds"),
					InitialDelaySeconds: viper.GetInt64("probes.redis.initialDelaySeconds"),
					FailureThreshold:    viper.GetInt64("probes.redis.failureThreshold"),
					SuccessThreshold:    viper.GetInt64("probes.redis.successThreshold"),
				},
			},
		}

		return &Config{
			cfg,
		}

	}
}

func setDefaults() {
	viper.SetDefault("server.host", defaults.Host)
	viper.SetDefault("server.port", defaults.Port)

	viper.SetDefault("logging.level", defaults.Verbosity)

	viper.SetDefault("probes.sqlite.timeoutSeconds", defaults.SqliteTimeout)
	viper.SetDefault("probes.sqlite.periodSeconds", defaults.SqlitePeriod)
	viper.SetDefault("probes.sqlite.initialDelaySeconds", defaults.SqliteInitialDelay)
	viper.SetDefault("probes.sqlite.failureThreshold", defaults.SqliteFailureThreshold)
	viper.SetDefault("probes.sqlite.successThreshold", defaults.SqliteSuccessThreshold)

	viper.SetDefault("probes.redis.timeoutSeconds", defaults.RedisTimeout)
	viper.SetDefault("probes.redis.periodSeconds", defaults.RedisPeriod)
	viper.SetDefault("probes.redis.initialDelaySeconds", defaults.RedisInitialDelay)
	viper.SetDefault("probes.redis.failureThreshold", defaults.RedisFailureThreshold)
	viper.SetDefault("probes.redis.successThreshold", defaults.RedisSuccessThreshold)

	viper.SetDefault("database.backend", "sqlite")
	viper.SetDefault("database.sqlite.path", "")

}

func bindEnv() {
	viper.BindEnv("server.host", "SERVER_HOST")
	viper.BindEnv("server.port", "SERVER_PORT")

	viper.BindEnv("logging.level", "LOGGING_LEVEL")

	viper.BindEnv("probes.sqlite.timeoutSeconds", "PROBES_SQLITE_TIMEOUT_SECONDS")
	viper.BindEnv("probes.sqlite.periodSeconds", "PROBES_SQLITE_PERIOD_SECONDS")
	viper.BindEnv("probes.sqlite.initialDelaySeconds", "PROBES_SQLITE_INITIAL_DELAY_SECONDS")
	viper.BindEnv("probes.sqlite.failureThreshold", "PROBES_SQLITE_FAILURE_THRESHOLD")
	viper.BindEnv("probes.sqlite.successThreshold", "PROBES_SQLITE_SUCCESS_THRESHOLD")

	viper.BindEnv("probes.redis.timeoutSeconds", "PROBES_REDIS_TIMEOUT_SECONDS")
	viper.BindEnv("probes.redis.periodSeconds", "PROBES_REDIS_PERIOD_SECONDS")
	viper.BindEnv("probes.redis.initialDelaySeconds", "PROBES_REDIS_INITIAL_DELAY_SECONDS")
	viper.BindEnv("probes.redis.failureThreshold", "PROBES_REDIS_FAILURE_THRESHOLD")
	viper.BindEnv("probes.redis.successThreshold", "PROBES_REDIS_SUCCESS_THRESHOLD")
}

func bindFlags(flags *pflag.FlagSet) {
	viper.BindPFlag("server.host", flags.Lookup("host"))
	viper.BindPFlag("server.port", flags.Lookup("port"))

	viper.BindPFlag("logging.level", flags.Lookup("v"))

	viper.BindPFlag("probes.sqlite.timeoutSeconds", flags.Lookup("sqlite-timeout"))
	viper.BindPFlag("probes.sqlite.periodSeconds", flags.Lookup("sqlite-period"))
	viper.BindPFlag("probes.sqlite.initialDelaySeconds", flags.Lookup("sqlite-initial-delay"))
	viper.BindPFlag("probes.sqlite.failureThreshold", flags.Lookup("sqlite-failure-threshold"))
	viper.BindPFlag("probes.sqlite.successThreshold", flags.Lookup("sqlite-success-threshold"))

	viper.BindPFlag("probes.redis.timeoutSeconds", flags.Lookup("redis-timeout"))
	viper.BindPFlag("probes.redis.periodSeconds", flags.Lookup("redis-period"))
	viper.BindPFlag("probes.redis.initialDelaySeconds", flags.Lookup("redis-initial-delay"))
	viper.BindPFlag("probes.redis.failureThreshold", flags.Lookup("redis-failure-threshold"))
	viper.BindPFlag("probes.redis.successThreshold", flags.Lookup("redis-success-threshold"))
}

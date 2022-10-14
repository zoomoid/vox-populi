package defaults

const (
	Host        = "0.0.0.0"
	Port uint16 = 8080

	Verbosity int = 0

	SqliteTimeout          int64 = 1
	SqlitePeriod           int64 = 10
	SqliteInitialDelay     int64 = 0
	SqliteFailureThreshold int64 = 3
	SqliteSuccessThreshold int64 = 1

	RedisTimeout          int64 = 1
	RedisPeriod           int64 = 10
	RedisInitialDelay     int64 = 0
	RedisFailureThreshold int64 = 3
	RedisSuccessThreshold int64 = 1
)

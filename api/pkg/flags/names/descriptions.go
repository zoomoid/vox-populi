package names

const (
	HostDescription = "Hostname for the server to listen"
	PortDescrition  = "Port to listen to"

	VerbosityDescription = "Logging level for logging volume"

	SqliteTimeoutDescription          = "Timeout duration in seconds for the health probes of the SQLite client"
	SqlitePeriodDescription           = "Period duration in seconds for the health probes of the SQLite client"
	SqliteInitialDelayDescription     = "Initial delay duraton in seconds for the health probes of the SQLite client"
	SqliteFailureThresholdDescription = "Threshold for failed attempts to connect to SQLite after which the probe is marked as failed"
	SqliteSuccessThresholdDescription = "Threshold for the successful attempts to connect to SQLite after which the probe is marked as successful"

	RedisTimeoutDescription          = "Timeout duration in seconds for the health probes of the Redis client"
	RedisPeriodDescription           = "Period duration in seconds for the health probes of the Redis client"
	RedisInitialDelayDescription     = "Initial delay duraton in seconds for the health probes of the Redis client"
	RedisFailureThresholdDescription = "Threshold for failed attempts to connect to Redis after which the probe is marked as failed"
	RedisSuccessThresholdDescription = "Threshold for the successful attempts to connect to Redis after which the probe is marked as successful"
)

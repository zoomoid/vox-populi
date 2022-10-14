package config

type Config struct {
	*Cfg
}

type Cfg struct {
	Logging  Logging  `json:"logging,omitempty"`
	Probes   Probes   `json:"probes,omitempty"`
	Server   Server   `json:"server,omitempty"`
	Database Database `json:"database,omitempty"`
	Cache    Cache    `json:"cache,omitempty"`
}

type Logging struct {
	Level int `json:"level,omitempty"`
}

type Server struct {
	Host string `json:"host,omitempty"`
	Port uint16 `json:"port,omitempty"`
}

type Database struct {
	Backend DatabaseBackend `json:"backend"`
	SQLite  SQLite          `json:"sqlite,omitempty"`
}

type DatabaseBackend string

const (
	DatabaseBackendSQLite   DatabaseBackend = "sqlite"
	DatabaseBackendPostgres DatabaseBackend = "postgres"
)

type SQLite struct {
	Path string `json:"path,omitempty"`
}

type Cache struct {
	Host     string `json:"host,omitempty"`
	Port     uint16 `json:"port,omitempty"`
	URL      string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DB       int    `json:"db,omitempty"`
}

type Probes struct {
	SQLite ProbeConfiguration `json:"sqlite"`
	Redis  ProbeConfiguration `json:"redis"`
}

type ProbeConfiguration struct {
	TimeoutSeconds      int64 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int64 `json:"periodSeconds,omitempty"`
	InitialDelaySeconds int64 `json:"initialDelaySeconds,omitempty"`
	FailureThreshold    int64 `json:"failureThreshold,omitempty"`
	SuccessThreshold    int64 `json:"successThreshold,omitempty"`
}

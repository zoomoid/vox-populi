package options

import "github.com/spf13/pflag"

type DatabaseConfiguration struct {
	Driver string
	Source string
}

type Configuration struct {
	Database DatabaseConfiguration
}

func AddFlags(cfg *Configuration, fs *pflag.FlagSet) {
	fs.StringVar(&cfg.Database.Driver, "db-driver", "postgres", "SQL database driver")
	fs.StringVar(&cfg.Database.Source, "db-host", "localhost:5432", "SQL database hostname and port")
}

func NewServerConfiguration() *Configuration {
	return &Configuration{
		Database: DatabaseConfiguration{
			Driver: "postgres",
			Source: "localhost:5432",
		},
	}
}

package db

import (
	"fmt"
	"log/slog"
	"os"

	"dario.cat/mergo"
	"github.com/caarlos0/env"
	"github.com/leometzger/timescale-cli/internal/config"
)

type ConnectionInfo struct {
	Host     string `env:"PGHOST" envDefault:"localhost"`
	Port     uint16 `env:"PGPORT" envDefault:"5432"`
	User     string `env:"PGUSER"`
	Pass     string `env:"PGPASS"`
	Database string `env:"PGDATABASE" envDefault:"tsdb"`
}

func (c *ConnectionInfo) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Database,
	)
}

func NewConnectionInfo(host string, port uint16, database string, user string, pass string) *ConnectionInfo {
	return &ConnectionInfo{
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     pass,
		Database: database,
	}
}

func loadConnectionInfoEnv() *ConnectionInfo {
	connectionInfo := &ConnectionInfo{}

	if err := env.Parse(connectionInfo); err != nil {
		slog.Error("unable to load database parameters config %v\n", err)
		os.Exit(1)
	}

	return connectionInfo
}

func LoadConnectionInfoWithConfigFile(configFile config.ConfigFile) *ConnectionInfo {
	connectionInfo := loadConnectionInfoEnv()
	mergo.Merge(&connectionInfo, configFile)
	return connectionInfo
}

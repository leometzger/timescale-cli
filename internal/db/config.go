package db

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/leometzger/timescale-cli/internal/config"
)

type ConnectionInfo struct {
	Host     string `env:"PGHOST" envDefault:"localhost"`
	Port     uint16 `env:"PGPORT" envDefault:"5432"`
	User     string `env:"PGUSER" envDefault:"postgres"`
	Database string `env:"PGDATABASE" envDefault:"postgres"`
	Password string `env:"PGPASS"`
}

func (c *ConnectionInfo) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

func NewConnectionInfo(
	host string,
	port uint16,
	database string,
	user string,
	pass string,
) *ConnectionInfo {
	return &ConnectionInfo{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		Database: database,
	}
}

func LoadConnectionInfoEnv() *ConnectionInfo {
	connectionInfo := &ConnectionInfo{}

	if err := env.Parse(connectionInfo); err != nil {
		slog.Error("unable to load database parameters config.", "cause", err)
		os.Exit(1)
	}

	return connectionInfo
}

func LoadConnectionInfoWithConfigFile(configFile *config.ConfigEnvironment) *ConnectionInfo {
	connectionInfo := LoadConnectionInfoEnv()

	// @TODO it could have a better logic here (reflection maybe?)
	if configFile.Database != "" {
		connectionInfo.Database = configFile.Database
	}

	if configFile.Host != "" {
		connectionInfo.Host = configFile.Host
	}

	if configFile.Port != 0 {
		connectionInfo.Port = configFile.Port
	}

	if configFile.User != "" {
		connectionInfo.User = configFile.User
	}

	if configFile.Password != "" {
		connectionInfo.Password = configFile.Password
	}

	return connectionInfo
}

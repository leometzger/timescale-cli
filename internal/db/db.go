package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v5"
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

func LoadConnectionInfoEnv() *ConnectionInfo {
	connectionInfo := &ConnectionInfo{}

	if err := env.Parse(connectionInfo); err != nil {
		slog.Error("unable to load database parameters config %v\n", err)
		os.Exit(1)
	}

	return connectionInfo
}

// creates a connection to postgres database
func Connect(connectionInfo *ConnectionInfo) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connectionInfo.String())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("connected to database sucessfully")
	return conn
}

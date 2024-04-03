package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// creates a pgx connection to postgres database
func Connect(connectionInfo *ConnectionInfo) PgxIface {
	conn, err := pgx.Connect(context.Background(), connectionInfo.String())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Debug("connected to database sucessfully")
	return conn
}

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	Close(context.Context) error
}

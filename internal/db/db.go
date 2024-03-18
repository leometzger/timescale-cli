package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

// creates a pgx connection to postgres database
func Connect(connectionInfo *ConnectionInfo) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connectionInfo.String())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("connected to database sucessfully")
	return conn
}

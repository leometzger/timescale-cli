package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/config"
)

func Connect(connectionInfo *config.ConnectionInfo) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connectionInfo.String())
	if err != nil {
		log.Fatalf("unable to connect to database %v\n", err)
	}
	return conn
}

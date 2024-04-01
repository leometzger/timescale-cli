package testlib

import (
	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/db"
)

func GetConnection() *pgx.Conn {
	info := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")
	conn := db.Connect(info)
	return conn
}

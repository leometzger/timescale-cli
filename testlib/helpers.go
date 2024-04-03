package testlib

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/leometzger/timescale-cli/internal/db"
)

func GetConnection() db.PgxIface {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	info := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")
	conn := db.Connect(info)

	return conn
}

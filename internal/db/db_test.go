package db_test

import (
	"context"
	"testing"

	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestIsAbleToConnectToTimescale(t *testing.T) {
	info := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")

	conn := db.Connect(info)
	defer conn.Close(context.Background())

	assert.NotNil(t, conn)
}

func TestIsAbleToQueryTimescale(t *testing.T) {
	info := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")

	conn := db.Connect(info)
	defer conn.Close(context.Background())

	result, err := conn.Query(context.Background(), "SELECT 1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestIsAbleToConfigureDBForTesting(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())

	assert.True(t, true)
}

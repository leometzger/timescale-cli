package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAbleToConnectToTimescale(t *testing.T) {
	pass := "5f4dcc3b5aa765d61d8327deb882cf99"
	info := NewConnectionInfo("localhost", 5432, "postgres", "postgres", pass)

	conn := Connect(info)
	defer conn.Close(context.Background())

	assert.NotNil(t, conn)
}

func TestIsAbleToQueryTimescale(t *testing.T) {
	info := NewConnectionInfo("localhost", 5432, "postgres", "postgres", "passsword")

	conn := Connect(info)
	defer conn.Close(context.Background())

	result, err := conn.Query(context.Background(), "SELECT 1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

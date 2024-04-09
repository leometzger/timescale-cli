package config_test

import (
	"path"
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadDevelopmentConfigCorrectly(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "development")

	require.Nil(t, err)
	assert.Equal(t, uint16(5432), c.Port)
	assert.Equal(t, "localhost", c.Host)
	assert.Equal(t, "datawarehouse", c.Database)
	assert.Equal(t, "postgres", c.User)
	assert.Equal(t, "password", c.Password)
}

func TestLoadStagingConfigCorrectly(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "staging")

	require.Nil(t, err)
	assert.Equal(t, uint16(5434), c.Port)
	assert.Equal(t, "localhost-s", c.Host)
	assert.Equal(t, "datawarehouse-s", c.Database)
	assert.Equal(t, "postgres-s", c.User)
	assert.Equal(t, "password-s", c.Password)
}

func TestLoadDefaultEnvironmentIfEmptyEnv(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "")

	require.Nil(t, err)
	assert.Equal(t, uint16(5434), c.Port)
	assert.Equal(t, "localhost-s", c.Host)
	assert.Equal(t, "datawarehouse-s", c.Database)
	assert.Equal(t, "postgres-s", c.User)
	assert.Equal(t, "password-s", c.Password)
}

func TestLoadInvalidConfig(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "invalid.yaml"), "staging")

	assert.Nil(t, c)
	assert.NotNil(t, err)
}

func TestLoadInexistentEnvironmentConfig(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "something")

	assert.Nil(t, c)
	assert.NotNil(t, err)
}

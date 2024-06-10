package config_test

import (
	"path"
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/stretchr/testify/require"
)

func TestLoadDevelopmentConfigCorrectly(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "development")

	require.Nil(t, err)
	require.Equal(t, uint16(5432), c.Port)
	require.Equal(t, "localhost", c.Host)
	require.Equal(t, "datawarehouse", c.Database)
	require.Equal(t, "postgres", c.User)
	require.Equal(t, "password", c.Password)
}

func TestLoadStagingConfigCorrectly(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "staging")

	require.Nil(t, err)
	require.Equal(t, uint16(5434), c.Port)
	require.Equal(t, "localhost-s", c.Host)
	require.Equal(t, "datawarehouse-s", c.Database)
	require.Equal(t, "postgres-s", c.User)
	require.Equal(t, "password-s", c.Password)
}

func TestLoadDefaultEnvironmentIfEmptyEnv(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "")

	require.Nil(t, err)
	require.Equal(t, uint16(5434), c.Port)
	require.Equal(t, "localhost-s", c.Host)
	require.Equal(t, "datawarehouse-s", c.Database)
	require.Equal(t, "postgres-s", c.User)
	require.Equal(t, "password-s", c.Password)
}

func TestLoadInvalidConfig(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "invalid.yaml"), "staging")

	require.Nil(t, c)
	require.NotNil(t, err)
}

func TestLoadInexistentEnvironmentConfig(t *testing.T) {
	c, err := config.LoadConfig(path.Join("testdata", "valid.yaml"), "something")

	require.Nil(t, c)
	require.NotNil(t, err)
}

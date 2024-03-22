package config

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDevelopmentConfigCorrectly(t *testing.T) {
	config, err := LoadConfig(path.Join("testdata", "valid.yaml"), "development")

	assert.Nil(t, err)
	assert.Equal(t, uint16(5432), config.Port)
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "datawarehouse", config.Database)
	assert.Equal(t, "postgres", config.User)
	assert.Equal(t, "password", config.Password)
}

func TestLoadStagingConfigCorrectly(t *testing.T) {
	config, err := LoadConfig(path.Join("testdata", "valid.yaml"), "staging")

	assert.Nil(t, err)
	assert.Equal(t, uint16(5434), config.Port)
	assert.Equal(t, "localhost-s", config.Host)
	assert.Equal(t, "datawarehouse-s", config.Database)
	assert.Equal(t, "postgres-s", config.User)
	assert.Equal(t, "password-s", config.Password)
}

func TestLoadInvalidConfig(t *testing.T) {
	config, err := LoadConfig(path.Join("testdata", "invalid.yaml"), "staging")

	assert.Nil(t, config)
	assert.NotNil(t, err)
}

func TestLoadInexistentEnvironmentConfig(t *testing.T) {
	config, err := LoadConfig(path.Join("testdata", "valid.yaml"), "something")

	assert.Nil(t, config)
	assert.NotNil(t, err)
}

func TestCreateNewConfigSucessfully(t *testing.T) {
	tmp := path.Join(os.TempDir(), "dw-config")
	err := os.RemoveAll(tmp)
	if err != nil {
		log.Fatal("Error setting up the environment to test")
	}

	err = os.Mkdir(tmp, os.ModePerm)
	if err != nil {
		log.Fatal("Error setting up the environment to test")
	}

	err = CreateConfig(tmp)

	assert.Nil(t, err)
	assert.FileExists(t, path.Join(tmp, DefaultConfigFileName))
}

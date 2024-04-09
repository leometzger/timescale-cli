package config_test

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ConfigEnvTestCase struct {
	EnvName string
	Env     config.ConfigEnvironment
}

func TestCreateNewConfigSucessfully(t *testing.T) {
	tmp := path.Join(os.TempDir(), "dw-config")
	configPath := path.Join(tmp, config.DefaultConfigFileName)

	err := os.Mkdir(tmp, os.ModePerm)
	defer os.RemoveAll(tmp)
	if err != nil {
		log.Fatal("Error setting up the environment to test")
	}

	tests := []ConfigEnvTestCase{
		{
			EnvName: "development",
			Env: config.ConfigEnvironment{
				Host:     "localhost",
				Port:     uint16(5555),
				Database: "postgres",
				User:     "postgres",
				Password: "password",
			},
		},
		{
			EnvName: "staging",
			Env: config.ConfigEnvironment{
				Host:     "db.staging.timescale.com",
				Port:     uint16(5433),
				Database: "timescale",
				User:     "postgres",
			},
		},
		{
			EnvName: "production",
			Env: config.ConfigEnvironment{
				Host:     "db.timescale.com",
				Port:     uint16(5432),
				Database: "tsdb",
				User:     "timescale-prod",
			},
		},
	}

	for _, test := range tests {
		err = config.CreateConfig(test.EnvName, &config.ConfigEnvironment{
			Host:     test.Env.Host,
			Database: test.Env.Database,
			Port:     test.Env.Port,
			User:     test.Env.User,
			Password: test.Env.Password,
		}, configPath)

		assert.Nil(t, err)
		assert.FileExists(t, path.Join(tmp, config.DefaultConfigFileName))
	}

	for _, test := range tests {
		env, err := config.LoadConfig(configPath, test.EnvName)

		require.Nil(t, err)
		assert.Equal(t, test.Env.Database, env.Database)
		assert.Equal(t, test.Env.Port, env.Port)
		assert.Equal(t, test.Env.User, env.User)
		assert.Equal(t, test.Env.Password, env.Password)
	}
}

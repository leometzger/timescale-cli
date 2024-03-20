package db

import (
	"os"
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func setFakeEnvsDifferenteFromDefault() {
	os.Setenv("PGHOST", "testing.postgres")
	os.Setenv("PGPORT", "5433")
	os.Setenv("PGUSER", "user1")
	os.Setenv("PGDATABASE", "tsdb")
	os.Setenv("PGPASS", "password1")
}

func resetEnvs() {
	os.Unsetenv("PGHOST")
	os.Unsetenv("PGPORT")
	os.Unsetenv("PGUSER")
	os.Unsetenv("PGUSER")
	os.Unsetenv("PGUSER")
}

func TestLoadConfigInfoWithSaneDefaults(t *testing.T) {
	resetEnvs()
	conf := LoadConnectionInfoEnv()

	assert.Equal(t, "localhost", conf.Host)
	assert.Equal(t, uint16(5432), conf.Port)
	assert.Equal(t, "postgres", conf.User)
	assert.Equal(t, "postgres", conf.Database)
}

func TestCouldLoadConfigInfoFromEnvironment(t *testing.T) {
	setFakeEnvsDifferenteFromDefault()

	conf := LoadConnectionInfoEnv()

	assert.Equal(t, "testing.postgres", conf.Host)
	assert.Equal(t, uint16(5433), conf.Port)
	assert.Equal(t, "user1", conf.User)
	assert.Equal(t, "tsdb", conf.Database)
	assert.Equal(t, "password1", conf.Password)
}

func TestShouldBeAbleToMergeWithConfigFile(t *testing.T) {
	setFakeEnvsDifferenteFromDefault()
	confFile := &config.ConfigFile{
		Host:     "localhost2",
		Port:     5435,
		User:     "user.conf",
		Database: "db.conf",
		Password: "pass.conf",
	}

	conf := LoadConnectionInfoWithConfigFile(confFile)

	assert.Equal(t, "localhost2", conf.Host)
	assert.Equal(t, uint16(5435), conf.Port)
	assert.Equal(t, "user.conf", conf.User)
	assert.Equal(t, "db.conf", conf.Database)
	assert.Equal(t, "pass.conf", conf.Password)
}

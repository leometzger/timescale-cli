package config

import (
	"log"
	"os"
	"path"
	"strconv"
)

type ConfigEnvironment struct {
	Name     string
	Host     string
	Database string
	Port     uint16
	User     string
	Password string
	Default  bool
}

type CliOptions struct {
	Env        string
	ConfigPath string
}

func NewCliOptions() *CliOptions {
	return &CliOptions{}
}

// Used to the test environment.
// This params would not be used for production
func DefaultConfig() *ConfigEnvironment {
	return &ConfigEnvironment{
		Name:     "default",
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "",
		Password: "",
		Default:  false,
	}
}

func GetDefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(home, ".tsctl", DefaultConfigFileName)
}

func (c *ConfigEnvironment) Headers() []string {
	return []string{
		"DEFAULT",
		"NAME",
		"HOST",
		"PORT",
		"DATABASE",
		"USER",
	}
}

func (c *ConfigEnvironment) Values() []string {
	defaultConfig := ""
	if c.Default {
		defaultConfig = "✓"
	}

	return []string{
		defaultConfig,
		c.Name,
		c.Host,
		strconv.FormatInt(int64(c.Port), 10),
		c.Database,
		c.User,
	}
}

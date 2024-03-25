package config

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const DefaultConfigFileName = ".ts-config.yaml"

type ConfigEnvironment struct {
	Host     string
	Database string
	Port     uint16
	User     string
	Password string
}

type CliOptions struct {
	Env        string
	ConfigPath string
	Verbose    bool
}

func NewCliOptions() *CliOptions {
	return &CliOptions{}
}

func LoadConfig(path string, env string) (*ConfigEnvironment, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		slog.Info("could not open config file, using default configuration")
		return DefaultConfig(), nil
	}

	conf := make(map[string]*ConfigEnvironment)
	err = yaml.Unmarshal(fileData, conf)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the config: %s", err)
	}

	if conf[env] == nil {
		return nil, fmt.Errorf("environment %s not found in the config %s", env, path)
	}

	return &ConfigEnvironment{
		Host:     conf[env].Host,
		Port:     conf[env].Port,
		User:     conf[env].User,
		Password: conf[env].Password,
		Database: conf[env].Database,
	}, nil
}

// creates a config into defualt file
func CreateConfig(envName string, env *ConfigEnvironment, configPath string) error {
	config := make(map[string]*ConfigEnvironment)
	config[envName] = env

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error converting the default configuration to YAML: %v", err)
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path.Dir(configPath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("could not create directory of config path %v", err)
		}

		f, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("could not access config file %v", err)
		}
		f.Close()
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("error while writing the config: %v", err)
	}

	return nil
}

// Used to the test environment. This params would not be used for production
func DefaultConfig() *ConfigEnvironment {
	return &ConfigEnvironment{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "",
		Password: "",
	}
}

func DefaultConfPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(home, DefaultConfigFileName)
}

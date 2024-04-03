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
}

func NewCliOptions() *CliOptions {
	return &CliOptions{}
}

func loadConfig(path string) (map[string]*ConfigEnvironment, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %s", err)
	}

	conf := make(map[string]*ConfigEnvironment)
	err = yaml.Unmarshal(fileData, conf)
	if err != nil {
		return conf, fmt.Errorf("error while parsing the config (yaml format): %s", err)
	}

	return conf, nil
}

func LoadConfig(path string, env string) (*ConfigEnvironment, error) {
	conf, err := loadConfig(path)
	if conf == nil {
		return nil, fmt.Errorf("could not find config file %v", err)
	}

	if err != nil {
		return nil, err
	}

	if conf[env] == nil {
		return nil, fmt.Errorf("could not find environment \"%v\"", env)
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
	conf, _ := loadConfig(configPath)
	if conf == nil {
		slog.Info("could not find config, creating a new one")
		conf = make(map[string]*ConfigEnvironment)
	}

	conf[envName] = env
	data, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Errorf("error converting the default configuration to YAML: %v", err)
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err := createConfigFile(configPath)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("error while writing the config: %v", err)
	}

	return nil
}

func createConfigFile(configPath string) error {
	err := os.MkdirAll(path.Dir(configPath), 0666)
	if err != nil {
		return fmt.Errorf("could not create directory of config path %v", err)
	}

	f, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("could not access config file %v", err)
	}

	err = f.Chmod(0666)
	if err != nil {
		return fmt.Errorf("could not change permissions of config file %v", err)
	}

	f.Close()
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

func GetDefaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(home, ".tsctl", DefaultConfigFileName)
}

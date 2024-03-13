package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const DefaultConfigFileName = ".ts-config.yaml"

type ConfigFile struct {
	Host     string
	Database string
	Port     int
	User     string
	Password string
}

type CliOptions struct {
	Env        string
	ConfigPath string
	Verbose    bool
}

func NewCliOptions(configPath string, verbose bool, env string) *CliOptions {
	return &CliOptions{
		ConfigPath: configPath,
		Verbose:    verbose,
		Env:        env,
	}
}

func LoadConfig(path string, env string) (*ConfigFile, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading the config file: %s", err)
	}

	conf := make(map[string]*ConfigFile)
	err = yaml.Unmarshal(fileData, conf)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the config: %s", err)
	}

	if conf[env] == nil {
		return nil, fmt.Errorf("environment %s not found in the config %s", env, path)
	}

	return &ConfigFile{
		Host:     conf[env].Host,
		Port:     conf[env].Port,
		User:     conf[env].User,
		Password: conf[env].Password,
		Database: conf[env].Database,
	}, nil
}

func CreateConfig(dest string) error {
	config := make(map[string]*ConfigFile)
	config["development"] = DefaultConfig()

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error converting the default configuration to YAML: %v", err)
	}

	err = os.WriteFile(path.Join(dest, DefaultConfigFileName), data, 0644)
	if err != nil {
		return fmt.Errorf("error while writing the config: %v", err)
	}

	return nil
}

// Used to the test environment. This params would not be used for production
func DefaultConfig() *ConfigFile {
	return &ConfigFile{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "password",
	}
}

func DefaultConfPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(home, DefaultConfigFileName)
}

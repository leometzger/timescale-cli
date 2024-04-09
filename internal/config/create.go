package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const DefaultConfigFileName = "config.yaml"

// creates a config into default file
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

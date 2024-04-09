package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

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

	for k, v := range conf {
		v.Name = k
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

	if env == "" {
		for k, v := range conf {
			if v.Default {
				env = k
				break
			}
		}
	}

	if conf[env] == nil {
		return nil, fmt.Errorf("could not find environment \"%v\"", env)
	}

	return &ConfigEnvironment{
		Name:     conf[env].Name,
		Host:     conf[env].Host,
		Port:     conf[env].Port,
		User:     conf[env].User,
		Password: conf[env].Password,
		Database: conf[env].Database,
	}, nil
}

func ListConfigs(path string) ([]*ConfigEnvironment, error) {
	conf, err := loadConfig(path)
	if err != nil {
		return nil, fmt.Errorf("could not find config file %v", err)
	}

	var confs []*ConfigEnvironment
	for _, v := range conf {
		confs = append(confs, v)
	}

	return confs, nil
}

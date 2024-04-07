package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type PayloadMapping struct {
	PathName     string `yaml:"pathName"`
	Project      string `yaml:"project"`
	GsmSecretRef string `yaml:"gsmSecretRef"`
}

type Config struct {
	PayloadMapping []PayloadMapping `yaml:"payloadMapping"`
}

func getConfigStruct(configName string) (Config, error) {
	var config Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func readApplicationData(applicationpath string, config Config) (*PayloadMapping, error) {
	for _, application := range config.PayloadMapping {
		if application.PathName == applicationpath {
			return &application, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("no payload mapping found for application: %s", applicationpath))
}

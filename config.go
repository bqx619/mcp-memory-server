package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigVector struct {
	Provider string `yaml:"provider"`
	URL      string `yaml:"url"`
	Token    string `yaml:"token"`
}

type ConfigHTTP struct {
	Port int `yaml:"port"`
}

type Config struct {
	Vector *ConfigVector `yaml:"vector"`
	HTTP   *ConfigHTTP   `yaml:"http"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// replace env variables
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

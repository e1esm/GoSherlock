package main

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type LookupService struct {
	Name     string `yaml:"name"`
	BaseLink string `yaml:"base_link"`
}

type Config struct {
	Services []LookupService `yaml:"services"`
}

func NewCfg(confFile string) (*Config, error) {
	b, err := os.ReadFile(confFile)
	if err != nil {
		return nil, fmt.Errorf("%w: error reading from file: %s", err, confFile)
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, fmt.Errorf("%w: error unmarshalling into cfg", err)
	}

	return cfg, nil
}

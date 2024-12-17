package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AdocConfig struct {
	WorkflowDirectory string `yaml:"workflow-directory"`
}

func ParseAdocConfig(cfg string) (AdocConfig, error) {
	fmt.Println(cfg)
	cfgBytes, err := os.ReadFile(cfg)
	if err != nil {
		return AdocConfig{}, err
	}

	adocConfig := AdocConfig{}
	if err := yaml.Unmarshal(cfgBytes, &adocConfig); err != nil {
		return AdocConfig{}, err
	}

	return adocConfig, nil
}

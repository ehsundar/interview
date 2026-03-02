package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ehsundar/interview.git/ratelimiter/internal/ratelimiter/config"
)

type Config struct {
	Rules map[string]config.Rule `json:"rules,omitempty"`
}

func LoadConfig() (Config, error) {
	var config Config
	configFile, err := os.Open("config.json")
	if err != nil {
		return config, err
	}

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configBytes, &config)
	return config, err
}

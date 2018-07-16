package config

import (
	"encoding/json"
	"io/ioutil"
)

var defaultFilePath = "/home/solidity/mining_config.json"

type Config struct {
	Coin           string `json:"coin"`
	URL            string `json:"url"`
	APIKey         string `json:"api_key"`
	SendgridAPIKey string `json:"sendgrid_api_key"`
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	if filePath == "" {
		filePath = defaultFilePath
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

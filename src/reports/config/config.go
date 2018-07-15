package config

import (
	"encoding/json"
	"io/ioutil"
)

var urlTemplate = "https://%s.miningpoolhub.com/index.php?page=api&action=%s&api_key=%s"
var defaultFilePath = "/home/solidity/mining_config.json"

type Config struct {
	Coin   string `json:"coin"`
	URL    string `json:"url"`
	APIKey string `json:"api_key"`
}

func LoadConfig(filePath string) (*Config, error) {
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

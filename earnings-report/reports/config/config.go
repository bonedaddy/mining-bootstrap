package config

import (
	"encoding/json"
	"io/ioutil"
)

var defaultFilePath = "/home/solidity/mining_config.json"

// Config holds our various configuration parameters
type Config struct {
	Coin           string `json:"coin"`
	URL            string `json:"url"`
	MPHAPIKey      string `json:"mph_api_key"`
	SendgridAPIKey string `json:"sendgrid_api_key"`
	DBURL          string `json:"db_url"`
	DBUser         string `json:"db_user"`
	DBPass         string `json:"db_pass"`
}

// LoadConfigFromFile is used to load our config from a file
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

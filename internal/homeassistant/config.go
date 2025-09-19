package homeassistant

import (
	"encoding/json"
	"os"
)

type Config struct {
	BaseUrl string `json:"base_url"`
	Token string `json:"token"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err;
	}

	return &config, nil;
}

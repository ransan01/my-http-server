package config

import (
	"encoding/json"
	"os"
)

func LoadConfigFromJsonFile(filename string) (*Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	config := NewConfig()
	err = json.NewDecoder(jsonFile).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

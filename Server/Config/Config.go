package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Host     string         `json:"host"`
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config

	// Read the content of index.js file (which is JSON)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading file: %v", err)
	}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return config, nil
}

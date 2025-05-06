package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

type RoomConfig struct {
	PasswordLength int `json:"passwordLength"`
}
type ServerConfig struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

func LoadServerConfig(filePath string) (ServerConfig, error) {
	var config ServerConfig

	// Read the content of index.js file (which is JSON)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading file: %v", err)
	}
	if config.Protocol == "" {
		config.Protocol = "http"
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return config, nil
}

func LoadRoomConfig(filePath string) (RoomConfig, error) {
	var config RoomConfig

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
	if config.PasswordLength < 1 {
		config.PasswordLength = 4
	}
	return config, nil
}

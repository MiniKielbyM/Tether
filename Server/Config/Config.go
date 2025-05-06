package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig `json:"Server"`
	Room   RoomConfig   `json:"Room"`
}

type RoomConfig struct {
	PasswordLength int `json:"passwordLength"`
	RoomsPerHost   int `json:"roomsPerHost"`
}

type ServerConfig struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

// Load the entire config (both Server and Room)
func LoadConfig(filePath string) (Config, error) {
	var config Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Apply defaults if necessary
	if config.Server.Protocol == "" {
		config.Server.Protocol = "http"
	}
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	if config.Room.PasswordLength < 1 {
		config.Room.PasswordLength = 1
	}
	if config.Room.RoomsPerHost < 1 {
		config.Room.RoomsPerHost = 1
	}

	return config, nil
}

package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig `json:"Server"`
	Room   RoomConfig   `json:"Room"`
	Api    ApiConfig    `json:"Api"`
}

type ServerConfig struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type RoomConfig struct {
	PasswordLength int `json:"passwordLength"`
	RoomsPerHost   int `json:"roomsPerHost"`
}

type ApiConfig struct {
	Dev DevConfig `json:"Dev"`
}

type DevConfig struct {
	Enabled bool `json:"enabled"`
	Port    int  `json:"port"`
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
	if config.Server.Port < 1 {
		config.Server.Port = 8080
	}
	if config.Room.PasswordLength < 1 {
		config.Room.PasswordLength = 1
	}
	if config.Room.RoomsPerHost < 1 {
		config.Room.RoomsPerHost = 1
	}
	if config.Api.Dev.Port < 1 {
		config.Api.Dev.Port = 8081
	}
	if config.Api.Dev.Port == config.Server.Port {
		config.Api.Dev.Port = config.Api.Dev.Port - 1
		if config.Api.Dev.Port < 1 {
			config.Api.Dev.Port = 8081
		}
		fmt.Printf("API port cannot be the same as server port, changing API port to %d\n", config.Api.Dev.Port)
	}

	return config, nil
}

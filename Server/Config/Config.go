package Config

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Config struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	for _, addr := range addrs {
		// Check the address type and skip loopback
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

func LoadConfig(filePath string) (Config, error) {
	var config Config

	// Read the content of index.js file (which is JSON)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("error reading file: %v", err)
	}
	if config.Host == "" {
		config.Host = getLocalIP()
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

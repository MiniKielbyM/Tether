package Core

import (
	"encoding/json"
	"fmt"
	"log"
)

func route(msg []byte) {
	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Printf("Error unmarshaling message: %v", err)
	}
	switch message.Type {
	case "join", "create", "close", "closeall":
		Room(msg)
	default:
		fmt.Printf("[Unknown type] %s\n", message.Type)
	}
}

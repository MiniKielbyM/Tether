package Core

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/MiniKielbyM/Tether/Server/Config"
	"github.com/google/uuid"
)

var Rooms = []RoomData{}

func createRoom(host string, passwordLength int) RoomData {
	var room RoomData
	room.Clients = make([]string, 0)
	room.Host = host
	room.RoomID = uuid.New().String()
	fmt.Printf("%s", generatePassword(passwordLength))
	room.Password = generatePassword(passwordLength)
	Rooms = append(Rooms, room)
	return room
}

func generatePassword(length int) string {
	password := ""
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password += string(charset[randomIndex])
	}
	return password
}

func Room(msg []byte) {
	config, _ := Config.LoadRoomConfig("room.json")
	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Printf("Error unmarshaling message: %v", err)
	}
	switch strings.ToLower(message.Type) {
	case "join":
		dataBytes, _ := json.Marshal(message.Data)
		var data JoinJson
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			log.Printf("Error unmarshaling join data: %v", err)
		} else {
			fmt.Printf("[Join] %v\n", data)
		}
	case "create":
		fmt.Print(createRoom(message.Sender, config.PasswordLength))
	default:
		fmt.Printf("[Unknown type] %s\n", message.Type)
	}
}

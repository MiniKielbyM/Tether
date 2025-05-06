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

func createRoom(host string, passwordLength int) (RoomData, error) {
	for _, room := range Rooms {
		if room.Host == host {
			return RoomData{}, fmt.Errorf("error creating room: Host has reached maximum amount of allowed simultaneous rooms")
		}
	}
	var room RoomData
	room.Clients = make([]string, 0)
	room.Host = host
	room.RoomID = uuid.New().String()
	room.Password = generatePassword(passwordLength)
	Rooms = append(Rooms, room)
	return room, nil
}

func closeRoom(host string, password string) {
	for i, room := range Rooms {
		if room.Host == host && room.Password == password{
			Rooms = append(Rooms[:i], Rooms[i+1:]...)
		}
	}
}

func generatePassword(length int) string {
	password := ""
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password += string(charset[randomIndex])
	}
	for _, room := range Rooms {
		if room.Password == password {
			generatePassword(length)
		}
	}
	return password
}

func Room(msg []byte) {
	config, _ := Config.LoadConfig("config.json")
	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Print("Error unmarshaling message: %v", err)
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
		fmt.Print(createRoom(message.Sender, config.Room.PasswordLength))
		fmt.Printf("\n")
	case "close":
		dataBytes, _ := json.Marshal(message.Data)
		var data CloseJson
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			log.Printf("Error unmarshaling join data: %v", err)
		} else {
			fmt.Printf("[Close] %v\n", data)
		}
		closeRoom(message.Sender, data.Password)
	default:
		fmt.Printf("[Unknown type] %s\n", message.Type)
	}
}

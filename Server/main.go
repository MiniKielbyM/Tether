package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MiniKielbyM/Tether/Server/Config" // Adjust the import path as necessary
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		fmt.Printf("Received: %s\n", msg)

		// Echo it back
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// Load configuration
	config, err := Config.LoadConfig("./Config.json") // Adjust the path as necessary
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Printf("Loaded config: %+v\n", config)
	http.HandleFunc("/ws", handleWS)
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

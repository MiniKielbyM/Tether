package Core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MiniKielbyM/Tether/Server/Config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var config, err = Config.LoadServerConfig("config.json")

// prevent blocking of the main thread
func startServer() {
	http.HandleFunc("/ws", handleWS)
	if err := http.ListenAndServe(":"+fmt.Sprint(config.Port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Client connected %s\n", conn.RemoteAddr())

	for {
		// Read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
		}
		message.Sender = conn.RemoteAddr().String()
		msg, _ = json.Marshal(message)
		route(msg)
		time.Sleep(10 * time.Millisecond)
	}
}

func StartServer() {
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	go startServer()        // Start the server in a goroutine
	time.Sleep(time.Second) // Give the server a second to start
	if config.Name == "" && config.Version == "" {
		fmt.Printf("Tether server started on port %s\n", fmt.Sprint(config.Port))
	} else if config.Name == "" && config.Version != "" {
		fmt.Printf("Tether server(v%s) started on port %s\n", config.Version, fmt.Sprint(config.Port))
	} else if config.Name != "" && config.Version == "" {
		fmt.Printf("Tether server %s started on port %s\n", config.Name, fmt.Sprint(config.Port))
	} else {
		fmt.Printf("Tether server %s(v%s) started on port %s\n", config.Name, config.Version, fmt.Sprint(config.Port))
	}
	select {} // Block forever
}

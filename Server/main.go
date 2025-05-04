package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MiniKielbyM/Tether/Server/Config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var config, err = Config.LoadConfig("config.json")

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

	for {
		// Read message
		typ, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		fmt.Printf("Received: %s\n", msg)

		// Echo it back
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
		}
		fmt.Printf("Message type: %d\n", typ)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	go startServer()        // Start the server in a goroutine
	time.Sleep(time.Second) // Give the server a second to start
	if config.Name == "" && config.Version == "" {
		fmt.Printf("Tether server started on %s://%s:%s\n", config.Protocol, config.Host, fmt.Sprint(config.Port))
	} else if config.Name == "" && config.Version != "" {
		fmt.Printf("Tether server(v%s) started on %s://%s:%s\n", config.Version, config.Protocol, config.Host, fmt.Sprint(config.Port))
	} else if config.Name != "" && config.Version == "" {
		fmt.Printf("Tether server %s started on %s://%s:%s\n", config.Name, config.Protocol, config.Host, fmt.Sprint(config.Port))
	} else {
		fmt.Printf("Tether server %s(v%s) started on %s://%s:%s\n", config.Name, config.Version, config.Protocol, config.Host, fmt.Sprint(config.Port))
	}
	select {} // Block forever
}

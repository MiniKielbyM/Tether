package Core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MiniKielbyM/Tether/Server/Config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var config, err = Config.LoadConfig("config.json")

// prevent blocking of the main thread
func startServer() {
	http.HandleFunc("/ws", handleWS)
	if err := http.ListenAndServe(":"+fmt.Sprint(config.Server.Port), nil); err != nil {
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
	log.Printf("Client connected %s\n", conn.RemoteAddr())

	for {
		// Read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) || strings.Contains(err.Error(), "unexpected EOF") || errors.Is(err, io.EOF) {
				log.Printf("client disconnected: %s\n", conn.RemoteAddr())
			} else {
				log.Printf("Error reading message: %v\n", err)
			}
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

func startupMsg() {
	if config.Server.Name == "" && config.Server.Version == "" {
		log.Printf("Tether server started on port %s\n", fmt.Sprint(config.Server.Port))
	} else if config.Server.Name == "" && config.Server.Version != "" {
		log.Printf("Tether server(v%s) started on port %s\n", config.Server.Version, fmt.Sprint(config.Server.Port))
	} else if config.Server.Name != "" && config.Server.Version == "" {
		log.Printf("Tether server %s started on port %s\n", config.Server.Name, fmt.Sprint(config.Server.Port))
	} else {
		log.Printf("Tether server %s(v%s) started on port %s\n", config.Server.Name, config.Server.Version, fmt.Sprint(config.Server.Port))
	}
}

func StartServer() {
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	go startServer()        // Start the server in a goroutine
	time.Sleep(time.Second) // Give the server a second to start
	startupMsg()
	select {} // Block forever
}

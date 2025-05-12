package Core

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
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
var clients = make(map[string]*websocket.Conn)

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
	clients[conn.RemoteAddr().String()] = conn
	SendToClient(conn.RemoteAddr().String(), "[-hello world-]")
	for {
		// Read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) || strings.Contains(err.Error(), "unexpected EOF") || errors.Is(err, io.EOF) {
				delete(clients, conn.RemoteAddr().String())
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

func StartWsServer() {
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	go startServer()        // Start the server in a goroutine
	time.Sleep(time.Second) // Give the server a second to start
	startupMsg()
}
func RenderPage(w http.ResponseWriter, tmplPath string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Println("Execution error:", err)
	}
}

func StartApiServer() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data := PageData{
				Name: config.Server.Name,
			}
			RenderPage(w, "Core/HTML/API.Html", data)
			time.Sleep(10 * time.Millisecond)
		})
		if err := http.ListenAndServe(":"+fmt.Sprint(config.Api.Dev.Port), nil); err != nil {
			log.Fatalf("HTML server failed to start: %v", err)
		}
	}()
	log.Printf("Developer API server started on port %s\n", fmt.Sprint(config.Api.Dev.Port))
}

func SendToClient(id string, message string) error {
	conn, ok := clients[id]
	if !ok {
		return fmt.Errorf("client %s not connected", id)
	}
	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}

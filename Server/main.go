package main

import (
	"github.com/MiniKielbyM/Tether/Server/Core"
)

func main() {
	go Core.StartServer() // Start the server in a goroutine
}

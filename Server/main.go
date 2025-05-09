package main

import (
	"github.com/MiniKielbyM/Tether/Server/Core"
)

func main() {
	Core.StartApiServer()
	Core.StartWsServer()
	select {} //block forever
}

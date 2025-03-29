package main

import (
	"golang/lib/core"
	"golang/lib/game_server/configuration"
	"golang/lib/game_server/server"
)

func main() {
	core.Initialize("Game server starting...")
	configuration := configuration.Configuration{
		Port: "7070",
	}
	var server = server.New(&configuration)
	server.Start()
}

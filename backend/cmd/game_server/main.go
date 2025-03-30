package main

import (
	"backend/lib/core"
	"backend/lib/database"
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/server"
)

func main() {
	core.Initialize("Game server starting...")
	serverSettings := core.Configuration{
		Port:             "7072",
		ConnectionString: "",
		UserAgent:        "backend/auth",
	}
	configuration := configuration.Configuration{
		ServerSettings: &serverSettings,
	}

	database := database.New(serverSettings.ConnectionString)

	var server = server.New(&configuration, database)
	server.Start()
}

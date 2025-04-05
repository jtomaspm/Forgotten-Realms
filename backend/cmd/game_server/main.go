package main

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/server"
	"backend/pkg/core"
	"backend/pkg/database"
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

	db := database.New(serverSettings.ConnectionString)
	defer db.Close()

	s := server.New(&configuration, db)
	s.Start()
}

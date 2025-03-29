package main

import (
	"backend/lib/auth_server/configuration"
	"backend/lib/auth_server/server"
	"backend/lib/core"
	"backend/lib/database"
)

func main() {
	core.Initialize("Auth server starting...")
	serverSettings := core.Configuration{
		Port:             "7070",
		ConnectionString: "",
	}
	configuration := configuration.Configuration{
		ServerSettings: &serverSettings,
	}

	database := database.New(serverSettings.ConnectionString)

	var server = server.New(&configuration, database)
	server.Start()
}

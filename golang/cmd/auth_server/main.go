package main

import (
	"golang/lib/auth_server/configuration"
	"golang/lib/auth_server/server"
	"golang/lib/core"
	"golang/lib/database"
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

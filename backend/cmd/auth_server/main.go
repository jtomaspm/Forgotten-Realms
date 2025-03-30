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
		UserAgent:        "backend/auth",
	}
	githubSettings := configuration.GitHub{
		ClientId:     "",
		ClientSecret: "",
		RedirectUri:  "",
		Source:       "GitHub",
	}
	configuration := configuration.Configuration{
		JwtSecret: "",
		GitHub:    &githubSettings,
		Server:    &serverSettings,
	}

	database := database.New(serverSettings.ConnectionString)

	var server = server.New(&configuration, database)
	server.Start()
}

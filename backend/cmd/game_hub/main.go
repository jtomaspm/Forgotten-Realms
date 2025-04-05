package main

import (
	"backend/lib/game_hub/configuration"
	"backend/lib/game_hub/server"
	"backend/pkg/core"
	"backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	core.Initialize("Hub server starting...")
	_ = godotenv.Load()
	envVars, err := core.GetEnv()
	if err != nil {
		log.Fatalln(err)
	}
	serverSettings := core.Configuration{
		Port:             "7070",
		ConnectionString: "postgres://postgres:123@localhost:5432/testdb",
		UserAgent:        "backend/game_hub",
	}
	dockerSettings := core.Docker{
		Auth:  envVars.DockerAuth,
		Hub:   envVars.DockerHub,
		Token: envVars.DockerToken,
	}
	configuration := configuration.Configuration{
		Docker:         &dockerSettings,
		ServerSettings: &serverSettings,
	}

	db := database.New(serverSettings.ConnectionString)
	defer db.Close()

	err = database.Migrate(db, "./migrations/game_server/")
	if err != nil {
		log.Fatalln(err)
	}

	s := server.New(&configuration, db)
	s.Start()
}

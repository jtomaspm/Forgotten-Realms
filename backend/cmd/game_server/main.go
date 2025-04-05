package main

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/server"
	"backend/pkg/core"
	"backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	core.Initialize("Game server starting...")
	_ = godotenv.Load()
	coreEnv, err := core.GetEnv()
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	dbConfig := database.Configuration{
		Host:     coreEnv.DbHost,
		Port:     coreEnv.DbPort,
		Username: coreEnv.DbUser,
		Password: coreEnv.DbPassword,
		Database: coreEnv.DbName,
	}
	serverSettings := core.Configuration{
		Port:      coreEnv.ServerPort,
		UserAgent: coreEnv.UserAgent,
		Database:  &dbConfig,
	}
	dockerSettings := core.Docker{
		Auth:  coreEnv.DockerAuth,
		Hub:   coreEnv.DockerHub,
		Token: coreEnv.DockerToken,
	}
	configuration := configuration.Configuration{
		Docker:         &dockerSettings,
		ServerSettings: &serverSettings,
	}

	db, err := database.Migrate(dbConfig, "./migrations/game_server/")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	s := server.New(&configuration, db)
	s.Start()
}

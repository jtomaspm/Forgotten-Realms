package main

import (
	"backend/lib/game_hub/configuration"
	"backend/lib/game_hub/server"
	"backend/pkg/core"
	"backend/pkg/core/models"
	"backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	core.Initialize("Hub server starting...")
	_ = godotenv.Load()
	coreEnv, err := models.GetEnv()
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
	serverSettings := models.Configuration{
		Port:      coreEnv.ServerPort,
		UserAgent: coreEnv.UserAgent,
		Database:  &dbConfig,
	}
	dockerSettings := models.Docker{
		Auth:  coreEnv.DockerAuth,
		Hub:   coreEnv.DockerHub,
		Token: coreEnv.DockerToken,
	}
	configuration := configuration.Configuration{
		Docker: &dockerSettings,
		Server: &serverSettings,
	}

	db, err := database.Migrate(dbConfig, "./migrations/game_hub/")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	s := server.New(&configuration, db)
	s.Start()
}

package main

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/configuration/realm_settings"
	"backend/lib/game_server/server"
	"backend/pkg/core"
	"backend/pkg/core/models"
	"backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	core.Initialize("Game server starting...")
	_ = godotenv.Load()
	coreEnv, err := models.GetEnv()
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	realmConfig, err := configuration.GetEnv()
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
		Realm:  realmConfig,
		Docker: &dockerSettings,
		Server: &serverSettings,
	}

	db, err := database.Migrate(dbConfig, "./migrations/game_server/")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = realm_settings.SyncRealmSettings("./realm_settings/", db)
	if err != nil {
		log.Fatalln(err)
	}

	s := server.New(&configuration, db)
	_, err = s.RegisterInHub()
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}

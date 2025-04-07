package main

import (
	"backend/lib/auth_server/configuration"
	"backend/lib/auth_server/server"
	"backend/pkg/core"
	"backend/pkg/core/models"
	"backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	core.Initialize("Auth server starting...")
	_ = godotenv.Load()
	envVars, err := configuration.GetEnv()
	if err != nil {
		log.Fatalln(err)
	}
	coreEnv, err := models.GetEnv()
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
	githubSettings := configuration.GitHub{
		ClientId:     envVars.GitHubClientId,
		ClientSecret: envVars.GitHubClientSecret,
		RedirectUri:  envVars.GitHubRedirectUri,
		Source:       "GitHub",
	}
	configuration := configuration.Configuration{
		JwtSecret: envVars.JwtSecret,
		GitHub:    &githubSettings,
		Server:    &serverSettings,
	}

	db, err := database.Migrate(dbConfig, "./migrations/auth_server/")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	s := server.New(&configuration, db)
	s.Start()
}

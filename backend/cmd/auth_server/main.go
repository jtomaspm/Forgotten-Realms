package main

import (
	"backend/lib/auth_server/configuration"
	"backend/lib/auth_server/server"
	"backend/pkg/core"
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
	serverSettings := core.Configuration{
		Port:             "7070",
		ConnectionString: "postgres://postgres:123@localhost:5432/testdb",
		UserAgent:        "backend/auth",
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

	db := database.New(serverSettings.ConnectionString)
	defer db.Close()

	err = database.Migrate(db, "./migrations/auth_server/")
	if err != nil {
		log.Fatalln(err)
	}

	s := server.New(&configuration, db)
	s.Start()
}

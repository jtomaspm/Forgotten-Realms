package models

import (
	"backend/pkg/database"
	"fmt"
	"os"
)

type Docker struct {
	Auth  string
	Hub   string
	Token string
}

type EnvVariables struct {
	ServerPort  string
	DbUser      string
	DbPassword  string
	DbName      string
	DbHost      string
	DbPort      string
	UserAgent   string
	DockerAuth  string
	DockerHub   string
	DockerToken string
}

func GetEnv() (*EnvVariables, error) {
	var variables EnvVariables
	variables.DockerAuth = os.Getenv("DOCKER_AUTH")
	variables.DockerHub = os.Getenv("DOCKER_HUB")
	variables.DockerToken = os.Getenv("DOCKER_TOKEN")
	variables.ServerPort = os.Getenv("SERVER_PORT")
	variables.DbUser = os.Getenv("DB_USER")
	variables.DbPassword = os.Getenv("DB_PASSWORD")
	variables.DbName = os.Getenv("DB_NAME")
	variables.DbHost = os.Getenv("DB_HOST")
	variables.DbPort = os.Getenv("DB_PORT")
	variables.UserAgent = os.Getenv("USER_AGENT")

	if variables.ServerPort == "" ||
		variables.DbUser == "" ||
		variables.DbPassword == "" ||
		variables.DbName == "" ||
		variables.DbHost == "" ||
		variables.UserAgent == "" ||
		variables.DockerAuth == "" ||
		variables.DockerHub == "" ||
		variables.DockerToken == "" {
		return &variables, fmt.Errorf("missing required env variables")
	}

	return &variables, nil
}

type Configuration struct {
	Port      string
	UserAgent string
	Database  *database.Configuration
}

package core

import (
	"fmt"
	"os"
)

type Docker struct {
	Auth  string
	Hub   string
	Token string
}

type EnvVariables struct {
	DockerAuth  string
	DockerHub   string
	DockerToken string
}

func GetEnv() (*EnvVariables, error) {
	var variables EnvVariables
	variables.DockerAuth = os.Getenv("DOCKER_AUTH")
	variables.DockerHub = os.Getenv("DOCKER_HUB")
	variables.DockerToken = os.Getenv("DOCKER_TOKEN")

	if variables.DockerAuth == "" ||
		variables.DockerHub == "" ||
		variables.DockerToken == "" {
		return &variables, fmt.Errorf("missing required env variables")
	}

	return &variables, nil
}

type Configuration struct {
	Port             string
	ConnectionString string
	UserAgent        string
}

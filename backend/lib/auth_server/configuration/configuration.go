package configuration

import (
	"backend/pkg/core"
	"fmt"
	"os"
)

type GitHub struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Source       string
}

type Configuration struct {
	JwtSecret string
	GitHub    *GitHub
	Server    *core.Configuration
}

type EnvVariables struct {
	JwtSecret          string
	GitHubClientId     string
	GitHubClientSecret string
	GitHubRedirectUri  string
}

func GetEnv() (*EnvVariables, error) {
	var variables EnvVariables
	variables.JwtSecret = os.Getenv("JWT_SECRET")
	variables.GitHubClientId = os.Getenv("GITHUB_CLIENT_ID")
	variables.GitHubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	variables.GitHubRedirectUri = os.Getenv("GITHUB_REDIRECT_URI")

	if variables.GitHubClientId == "" ||
		variables.GitHubClientSecret == "" ||
		variables.JwtSecret == "" ||
		variables.GitHubRedirectUri == "" {
		return &variables, fmt.Errorf("missing required env variables")
	}

	return &variables, nil
}

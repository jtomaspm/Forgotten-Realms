package configuration

import "backend/lib/core"

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

package configuration

import "backend/pkg/core"

type Configuration struct {
	Docker         *core.Docker
	ServerSettings *core.Configuration
}

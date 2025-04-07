package configuration

import (
	"backend/pkg/core/models"
)

type Configuration struct {
	Docker *models.Docker
	Server *models.Configuration
}

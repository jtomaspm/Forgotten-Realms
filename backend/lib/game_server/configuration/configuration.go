package configuration

import (
	"backend/pkg/core/models"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const CHUNK_SIZE = 10
const MAP_SIZE = 1000

type GameServerConfiguration struct {
	Name           string
	PublicEndpoint string
	Id             uuid.UUID
}

func GetEnv() (*GameServerConfiguration, error) {
	var variables GameServerConfiguration
	variables.Name = os.Getenv("REALM_NAME")
	variables.PublicEndpoint = os.Getenv("PUBLIC_ENDPOINT")

	if variables.Name == "" ||
		variables.PublicEndpoint == "" {
		return &variables, fmt.Errorf("missing required env variables")
	}

	return &variables, nil
}

type Configuration struct {
	Realm  *GameServerConfiguration
	Docker *models.Docker
	Server *models.Configuration
}

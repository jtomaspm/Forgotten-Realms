package views

import (
	"backend/lib/auth_server/dal/models"

	"github.com/google/uuid"
)

type AccountShort struct {
	Id    uuid.UUID
	Name  string
	Email string
	Role  models.Role
}

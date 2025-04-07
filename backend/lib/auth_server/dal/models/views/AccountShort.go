package views

import (
	"backend/pkg/core/models"

	"github.com/google/uuid"
)

type AccountShort struct {
	Id    uuid.UUID
	Name  string
	Email string
	Role  models.Role
}

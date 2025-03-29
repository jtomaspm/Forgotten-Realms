package models

import (
	"time"

	"github.com/google/uuid"
)

type World struct {
	Id        uuid.UUID
	Name      string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

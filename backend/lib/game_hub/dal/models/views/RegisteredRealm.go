package views

import (
	"time"

	"github.com/google/uuid"
)

type RegisteredRealm struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Registered bool      `json:"registered"`
	CreatedAt  time.Time `json:"created_at"`
}

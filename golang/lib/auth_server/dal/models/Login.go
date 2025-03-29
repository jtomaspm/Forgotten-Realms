package models

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	Id        uuid.UUID
	AccountId uuid.UUID
	CreatedAt time.Time
	IpAddress string
}

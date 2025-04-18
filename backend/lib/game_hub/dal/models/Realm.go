package models

import (
	"time"

	"github.com/google/uuid"
)

type Realm struct {
	Id        uuid.UUID
	Name      string
	Api       string
	Status    RealmStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

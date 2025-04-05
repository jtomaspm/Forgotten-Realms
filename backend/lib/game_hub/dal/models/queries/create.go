package queries

import "github.com/google/uuid"

type CreateRealm struct {
	Name string
	Api  string
}

type CreateAccountRealm struct {
	AccountId uuid.UUID
	RealmId   uuid.UUID
}

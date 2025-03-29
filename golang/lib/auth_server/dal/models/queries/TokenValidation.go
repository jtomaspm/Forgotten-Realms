package queries

import "golang/lib/auth_server/dal/models"

type TokenValidation struct {
	Account models.Account
	Session models.Session
}

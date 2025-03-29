package queries

import "backend/lib/auth_server/dal/models"

type AccountDetails struct {
	Account           models.Account
	AccountProperties models.AccountProperties
	LastLogin         models.Login
	Session           models.Session
	Worlds            []models.World
}

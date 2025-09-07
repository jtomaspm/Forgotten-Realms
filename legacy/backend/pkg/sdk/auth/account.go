package auth

import (
	"backend/pkg/core/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Account struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

func (account *Account) IsAuthorized(requiredRole models.Role) bool {
	r, err := models.FromString(account.Role)
	if err != nil {
		return false
	}
	return r <= requiredRole
}

type GetAccountRequest struct {
	Auth  string
	Token string
}

func GetAccount(request *GetAccountRequest) (Account, error) {
	var account Account
	reqPath := "http://" + request.Auth + "/api/account?token=" + request.Token

	resp, err := http.Get(reqPath)
	if err != nil {
		return account, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return account, fmt.Errorf("failed to get account: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

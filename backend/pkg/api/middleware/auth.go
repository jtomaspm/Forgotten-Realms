package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Account struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

func getAuthToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is required")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization format")
	}
	return parts[1], nil
}

func getAccount(ctx *gin.Context, authServer string) (Account, error) {
	var account Account
	token, err := getAuthToken(ctx)
	if err != nil {
		return account, err
	}
	reqPath := authServer + "/api/account?token=" + token

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

func AuthMiddleware(authServer string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account, err := getAccount(ctx, authServer)
		if err == nil {
			ctx.Set("account", account)
			ctx.Next()
			return
		}
	}
}

func GetAccountFromContext(ctx *gin.Context) (Account, error) {
	account, exists := ctx.Get("account")
	if !exists {
		return Account{}, fmt.Errorf("account not found in context")
	}
	acc, ok := account.(Account)
	if !ok {
		return Account{}, fmt.Errorf("account is not of type Account")
	}
	return acc, nil
}

func IsRequestInternal(ctx *gin.Context, validationToken string) bool {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return false
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Internal" {
		return false
	}
	return parts[1] == validationToken
}

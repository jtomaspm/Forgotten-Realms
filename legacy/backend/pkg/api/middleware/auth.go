package middleware

import (
	"backend/pkg/sdk/auth"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

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

func getAccount(ctx *gin.Context, authServer string) (auth.Account, error) {
	token, err := getAuthToken(ctx)
	if err != nil {
		return auth.Account{}, err
	}
	return auth.GetAccount(&auth.GetAccountRequest{
		Auth:  authServer,
		Token: token,
	})
}

func AuthMiddleware(authServer string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account, err := getAccount(ctx, authServer)
		if err == nil {
			ctx.Set("account", account)
		}
		ctx.Next()
	}
}

func GetAccountFromContext(ctx *gin.Context) (auth.Account, error) {
	account, exists := ctx.Get("account")
	if !exists {
		return auth.Account{}, fmt.Errorf("account not found in context")
	}
	acc, ok := account.(auth.Account)
	if !ok {
		return auth.Account{}, fmt.Errorf("account is not of type Account")
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

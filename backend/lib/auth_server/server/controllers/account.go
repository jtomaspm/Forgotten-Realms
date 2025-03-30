package controllers

import (
	"backend/lib/auth_server/configuration"
	dalModels "backend/lib/auth_server/dal/models"
	"backend/lib/auth_server/dal/models/queries"
	"backend/lib/auth_server/dal/services/account"
	"backend/lib/auth_server/server/models"
	"backend/lib/database"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AccountController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *AccountController) Mount(basePath string, engine *gin.Engine) {
	path := basePath + "/account"
	engine.POST(path+"/create", controller.create)
}

func (controller *AccountController) create(ctx *gin.Context) {
	name := ctx.GetHeader("Name")
	if name == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Name header is required"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		return
	}
	tokenString := parts[1]
	var claims models.NewUserClaims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return controller.Configuration.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	verificationToken, err := account.Create(ctx, controller.Database.Pool, &queries.CreateAccount{
		ExternalId: claims.ExternalId,
		Source:     claims.Source,
		Email:      claims.Email,
		Role:       dalModels.PLAYER,
		Name:       name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new account"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account created, verify email using token",
		"token":   verificationToken,
	})
}

package controllers

import (
	"backend/lib/auth_server/configuration"
	dalModels "backend/lib/auth_server/dal/models"
	"backend/lib/auth_server/dal/models/queries"
	"backend/lib/auth_server/dal/services/account"
	"backend/lib/auth_server/dal/services/session"
	"backend/lib/auth_server/server/models"
	"backend/pkg/database"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AccountController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *AccountController) Mount(basePath string, engine *gin.Engine) {
	path := basePath + "/account"
	engine.POST(path+"", controller.create)
	engine.GET(path+"", controller.get)
	engine.GET(path+"/id", controller.getId)
	engine.GET(path+"/verify", controller.verifyEmail)
}

func (controller *AccountController) get(ctx *gin.Context) {
	errorStatus := http.StatusBadRequest
	tokenStr := ctx.Query("token")
	if tokenStr == "" {
		ctx.JSON(errorStatus, gin.H{"error": "Missing token parameter"})
		return
	}
	token, err := uuid.Parse(tokenStr)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": "Invalid token"})
		return
	}
	account, err := session.GetAccount(ctx, controller.Database.Pool, token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":    account.Id,
		"name":  account.Name,
		"email": account.Email,
		"role":  account.Role.String(),
	})
}

func (controller *AccountController) getId(ctx *gin.Context) {
	errorStatus := http.StatusBadRequest
	tokenStr := ctx.Query("token")
	if tokenStr == "" {
		ctx.JSON(errorStatus, gin.H{"error": "Missing token parameter"})
		return
	}
	token, err := uuid.Parse(tokenStr)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": "Invalid token"})
		return
	}
	id, err := session.GetAccountId(ctx, controller.Database.Pool, token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (controller *AccountController) create(ctx *gin.Context) {
	var body struct {
		Name                   string `json:"name"`
		SendEmailNotifications bool   `json:"send_email_notifications"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if body.Name == "" {
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
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(controller.Configuration.JwtSecret), nil
	})
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	verificationToken, err := account.Create(ctx, controller.Database.Pool, &queries.CreateAccount{
		ExternalId:             claims.ExternalId,
		Source:                 claims.Source,
		Email:                  claims.Email,
		Role:                   dalModels.PLAYER,
		Name:                   body.Name,
		SendEmailNotifications: body.SendEmailNotifications,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new account"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Account created, verify email using token",
		"token":   verificationToken,
	})
}

func (controller *AccountController) verifyEmail(ctx *gin.Context) {
	errorStatus := http.StatusBadRequest
	tokenStr := ctx.Query("token")
	if tokenStr == "" {
		ctx.JSON(errorStatus, gin.H{"error": "Missing token parameter"})
		return
	}
	token, err := uuid.Parse(tokenStr)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": "Invalid token"})
		return
	}
	err = account.Verify(ctx, controller.Database.Pool, token)
	if err != nil {
		ctx.JSON(errorStatus, gin.H{"error": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Success, you may now login",
	})
}

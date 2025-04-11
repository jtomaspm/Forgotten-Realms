package controllers

import (
	"backend/lib/game_hub/configuration"
	"backend/lib/game_hub/dal/models"
	"backend/lib/game_hub/dal/models/queries"
	"backend/lib/game_hub/dal/models/views"
	"backend/lib/game_hub/dal/services/realms"
	"backend/pkg/api/middleware"
	"backend/pkg/database"
	sdkRealms "backend/pkg/sdk/hub/realms"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RealmsController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *RealmsController) Mount(basePath string, engine *gin.Engine) {
	engine.GET(basePath+"/realm", controller.getRealms)
	engine.GET(basePath+"/realm/:id", controller.getRealm)
	engine.POST(basePath+"/realm", controller.registerRealm)
	engine.POST(basePath+"/realm/account", controller.registerAccount)
}

func (Controller *RealmsController) getRealms(ctx *gin.Context) {
	var result []views.RegisteredRealm
	acc, err := middleware.GetAccountFromContext(ctx)
	if err != nil {
		result, err = realms.GetAll(ctx, Controller.Database.Pool)
	} else {
		result, err = realms.GetByAccountId(ctx, Controller.Database.Pool, acc.Id)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get realms"})
		log.Println("Failed to get realms:", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"realms": result})
}

func (Controller *RealmsController) getRealm(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid realm id"})
		log.Println("Failed to get realm:", err)
		return
	}
	var result models.Realm
	result, err = realms.GetById(ctx, Controller.Database.Pool, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get realm"})
		log.Println("Failed to get realms:", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":         result.Id,
		"name":       result.Name,
		"api":        result.Api,
		"status":     result.Status.String(),
		"created_at": result.CreatedAt,
		"updated_at": result.UpdatedAt,
	})
}

func (controller *RealmsController) registerRealm(ctx *gin.Context) {
	if !middleware.IsRequestInternal(ctx, controller.Configuration.Docker.Token) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var realm queries.CreateRealm
	if err := ctx.ShouldBindJSON(&realm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	realmId, err := realms.RegisterRealm(ctx, controller.Database.Pool, &realm)
	if err != nil {
		realmId, err = realms.GetByCreateQuery(ctx, controller.Database.Pool, &realm)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register realm"})
		log.Println("Failed to register realm:", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": realmId})
}

func (controller *RealmsController) registerAccount(ctx *gin.Context) {
	if !middleware.IsRequestInternal(ctx, controller.Configuration.Docker.Token) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var command sdkRealms.RegisterAccountRequestBody
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := realms.RegisterAccount(ctx, controller.Database.Pool, &command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register account"})
		log.Println("Failed to register account:", err)
		return
	}
	ctx.Status(http.StatusCreated)
}

package controllers

import (
	"backend/lib/game_hub/configuration"
	"backend/lib/game_hub/dal/models/queries"
	"backend/lib/game_hub/dal/models/views"
	"backend/lib/game_hub/dal/services/realms"
	"backend/pkg/api/middleware"
	"backend/pkg/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RealmsController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *RealmsController) Mount(basePath string, engine *gin.Engine) {
	engine.GET(basePath+"/realm", controller.getRealms)
	engine.POST(basePath+"/realm", controller.registerRealm)
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

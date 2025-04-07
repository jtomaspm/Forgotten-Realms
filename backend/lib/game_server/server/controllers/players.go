package controllers

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/dal/models/queries"
	"backend/lib/game_server/dal/services/players"
	"backend/pkg/api/middleware"
	"backend/pkg/core/models"
	"backend/pkg/database"
	"backend/pkg/sdk/hub/realms"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayersController struct {
	Configuration *configuration.Configuration
	Database      *database.Database
}

func (controller *PlayersController) Mount(basePath string, engine *gin.Engine) {
	engine.POST(basePath+"/player", controller.create)
	engine.GET(basePath+"/player/:id", controller.getById)
	engine.GET(basePath+"/player", controller.query)
}

func (controller *PlayersController) getById(ctx *gin.Context) {

}

func (controller *PlayersController) query(ctx *gin.Context) {

}

func (controller *PlayersController) create(ctx *gin.Context) {
	account, err := middleware.GetAccountFromContext(ctx)
	if err != nil || !account.IsAuthorized(models.PLAYER) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	var command queries.CreatePlayer
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	transaction, err := controller.Database.Pool.Begin(ctx)
	defer transaction.Rollback(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player"})
		return
	}
	id, err := players.Create(ctx, transaction, &command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player"})
		return
	}
	err = realms.RegisterAccount(&realms.RegisterAccountRequest{
		Hub:           controller.Configuration.Docker.Hub,
		InternalToken: controller.Configuration.Docker.Token,
		Body: &realms.RegisterAccountRequestBody{
			AccountId: account.Id,
			RealmId:   controller.Configuration.Realm.Id,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player in hub"})
		return
	}
	err = transaction.Commit(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

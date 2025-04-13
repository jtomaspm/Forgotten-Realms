package controllers

import (
	"backend/lib/game_server/configuration"
	dalModels "backend/lib/game_server/dal/models"
	"backend/lib/game_server/dal/models/queries"
	"backend/lib/game_server/dal/services/players"
	"backend/lib/game_server/dal/services/settings"
	"backend/lib/game_server/services/village_s"
	"backend/pkg/api/middleware"
	"backend/pkg/core/models"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"backend/pkg/sdk/hub/realms"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player id"})
		return
	}
	var result dalModels.Player
	result, err = players.GetById(ctx, controller.Database.Pool, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":         result.Id,
		"faction":    result.Faction.String(),
		"created_at": result.CreatedAt,
		"updated_at": result.UpdatedAt,
	})
}

func (controller *PlayersController) query(ctx *gin.Context) {

}

func (controller *PlayersController) create(ctx *gin.Context) {
	account, err := middleware.GetAccountFromContext(ctx)
	if err != nil || !account.IsAuthorized(models.PLAYER) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	open, err := realms.IsRealmOpen(&realms.GetRealmRequest{
		Hub:     controller.Configuration.Docker.Hub,
		RealmId: controller.Configuration.Realm.Id,
	})
	if err != nil || !open {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Server closed."})
		log.Println(err)
		return
	}
	var playerProps struct {
		Faction  string `json:"faction"`
		Location string `json:"location"`
	}
	if err := ctx.ShouldBindJSON(&playerProps); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	faction, err := enum.FactionFromString(playerProps.Faction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid faction"})
		return
	}
	command := queries.CreatePlayer{
		AccountId: account.Id,
		Faction:   faction,
	}
	transaction, err := controller.Database.Pool.Begin(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player"})
		return
	}
	defer transaction.Rollback(ctx)
	id, err := players.Create(ctx, transaction, &command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating player"})
		return
	}

	realmSettings, err := settings.GetRealmSettings(ctx, transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating village"})
		return
	}
	location, err := enum.SpawnLocationFromString(playerProps.Location)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location"})
		return
	}
	playerVillage, err := village_s.SpawnVillage(ctx, transaction, account.Id, location, realmSettings)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating village"})
		return
	}

	//Do after all operations in game database
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
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
		"first_village": gin.H{
			"coord_x": playerVillage.CoordX,
			"coord_y": playerVillage.CoordX,
		},
	})
}

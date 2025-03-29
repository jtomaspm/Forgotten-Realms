package server

import (
	"golang/lib/game_server/configuration"
	"golang/lib/game_server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	basePath string
	Engine   *gin.Engine
}

func NewRouter(basePath string, configuration *configuration.Configuration) *Router {
	engine := gin.Default()

	mount_TestController(basePath, configuration, engine)

	return &Router{
		Engine: engine,
	}
}

func mount_TestController(basePath string, configuration *configuration.Configuration, engine *gin.Engine) {
	c := controllers.TestController{
		Configuration: configuration,
	}
	c.Mount(basePath, engine)
}

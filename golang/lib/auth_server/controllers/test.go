package controllers

import (
	"golang/lib/auth_server/configuration"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	Configuration *configuration.Configuration
}

func (controller *TestController) Mount(basePath string, engine *gin.Engine) {
	engine.GET(basePath+"/test", controller.test)
}

func (Controller *TestController) test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"test": "message",
	})
}

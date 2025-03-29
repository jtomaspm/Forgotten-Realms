package api

import "github.com/gin-gonic/gin"

type Controller interface {
	Mount(basePath string, engine *gin.Engine)
}

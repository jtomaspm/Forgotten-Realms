package api

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	BasePath    string
	Controllers []Controller
}

type Router struct {
	Engine *gin.Engine
}

func NewRouter(routes []Route) *Router {
	engine := gin.Default()

	for _, route := range routes {
		for _, controller := range route.Controllers {
			controller.Mount(route.BasePath, engine)
		}
	}

	return &Router{
		Engine: engine,
	}
}

package api

import (
	"backend/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

type Route struct {
	BasePath    string
	Controllers []Controller
}

type Router struct {
	Engine *gin.Engine
}

type AuthSettings struct {
	AuthServer string
	UseAuth    bool
}

func NewRouter(routes []Route, auth *AuthSettings) *Router {
	engine := gin.Default()

	if auth.UseAuth {
		engine.Use(middleware.AuthMiddleware(auth.AuthServer))
	}

	for _, route := range routes {
		for _, controller := range route.Controllers {
			controller.Mount(route.BasePath, engine)
		}
	}

	return &Router{
		Engine: engine,
	}
}

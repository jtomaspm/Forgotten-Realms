package server

import (
	"backend/lib/game_hub/configuration"
	"backend/lib/game_hub/server/controllers"
	"backend/pkg/api"
	"backend/pkg/api/middleware"
	"backend/pkg/database"
	"net/http"
)

type Server struct {
	Configuration *configuration.Configuration
	router        *api.Router
	http          *http.Server
	database      *database.Database
}

func New(configuration *configuration.Configuration, database *database.Database) *Server {
	var routes = []api.Route{
		{
			BasePath: "/api",
			Controllers: []api.Controller{
				&controllers.RealmsController{
					Configuration: configuration,
					Database:      database,
				},
			},
		},
	}
	router := api.NewRouter(routes)
	router.Engine.Use(middleware.AuthMiddleware(configuration.Docker.Auth))
	server := &http.Server{
		Addr:    ":" + configuration.ServerSettings.Port,
		Handler: router.Engine,
	}
	return &Server{
		Configuration: configuration,
		router:        router,
		http:          server,
		database:      database,
	}
}

func (server *Server) Start() {
	server.http.ListenAndServe()
}

package server

import (
	"backend/lib/core/api"
	"backend/lib/database"
	"backend/lib/game_server/configuration"
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
			BasePath:    "/api",
			Controllers: []api.Controller{},
		},
	}
	router := api.NewRouter(routes)
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

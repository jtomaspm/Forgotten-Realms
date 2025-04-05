package server

import (
	"backend/lib/game_server/configuration"
	"backend/pkg/api"
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
			BasePath:    "/api",
			Controllers: []api.Controller{},
		},
	}
	router := api.NewRouter(routes, &api.AuthSettings{
		AuthServer: configuration.Docker.Auth,
		UseAuth:    true,
	})
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

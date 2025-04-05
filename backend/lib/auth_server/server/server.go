package server

import (
	"backend/lib/auth_server/configuration"
	"backend/lib/auth_server/server/controllers"
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
			BasePath: "/api",
			Controllers: []api.Controller{
				&controllers.GithubController{Configuration: configuration, Database: database},
				&controllers.AccountController{Configuration: configuration, Database: database},
			},
		},
	}
	router := api.NewRouter(routes)
	server := &http.Server{
		Addr:    ":" + configuration.Server.Port,
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

package server

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/server/controllers"
	"backend/pkg/api"
	"backend/pkg/database"
	"backend/pkg/sdk/hub/realms"
	"net/http"

	"github.com/google/uuid"
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
				&controllers.PlayersController{
					Configuration: configuration,
					Database:      database,
				},
			},
		},
	}
	router := api.NewRouter(routes, &api.AuthSettings{
		AuthServer: configuration.Docker.Auth,
		UseAuth:    true,
	})
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

func (server *Server) RegisterInHub() (uuid.UUID, error) {
	id, err := realms.RegisterRealm(&realms.RegisterRealmRequest{
		Hub:           server.Configuration.Docker.Hub,
		InternalToken: server.Configuration.Docker.Token,
		Body: &realms.RegisterRealmRequestBody{
			Name: server.Configuration.Realm.Name,
			Api:  server.Configuration.Realm.PublicEndpoint,
		},
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	server.Configuration.Realm.Id = id
	return id, nil
}

func (server *Server) Start() {
	server.http.ListenAndServe()
}

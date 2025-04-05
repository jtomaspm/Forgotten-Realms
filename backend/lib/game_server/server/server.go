package server

import (
	"backend/lib/game_server/configuration"
	"backend/pkg/api"
	"backend/pkg/database"
	"bytes"
	"encoding/json"
	"fmt"
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

func (server *Server) RegisterInHub() (uuid.UUID, error) {
	var response struct {
		Id uuid.UUID `json:"id"`
	}
	reqPath := "http://" + server.Configuration.Docker.Hub + "/api/realm"

	body := struct {
		Name string `json:"name"`
		Api  string `json:"api"`
	}{
		Name: server.Configuration.ServerSettings.UserAgent,
		Api:  server.Configuration.ServerSettings.UserAgent + ":" + server.Configuration.ServerSettings.Port,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return response.Id, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", reqPath, bytes.NewReader(bodyBytes))
	if err != nil {
		return response.Id, err
	}

	req.Header.Set("Authorization", "Internal "+server.Configuration.Docker.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response.Id, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return response.Id, fmt.Errorf("failed to register realm: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Id, err
	}
	return response.Id, nil
}

func (server *Server) Start() {
	server.http.ListenAndServe()
}

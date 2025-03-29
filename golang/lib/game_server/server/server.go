package server

import (
	"golang/lib/game_server/configuration"
	"net/http"
)

type Server struct {
	Configuration *configuration.Configuration
	router        *Router
	http          *http.Server
}

func New(configuration *configuration.Configuration) *Server {
	router := NewRouter("/api", configuration)
	server := &http.Server{
		Addr:    ":" + configuration.Port,
		Handler: router.Engine,
	}
	return &Server{
		Configuration: configuration,
		router:        router,
		http:          server,
	}
}

func (server *Server) Start() {
	server.http.ListenAndServe()
}

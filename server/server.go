package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/database"
	"time"
)

type Config struct {
	Database *database.Config `yaml:"db"`
}

type Server struct {
	Database *database.Database

	server *http.Server
}

func CreateServer(config *Config) (*Server, error) {
	server := Server{}
	var err error

	// Init other services
	server.Database, err = database.Init(config.Database)
	if err != nil {
		return nil, err
	}

	// Init server
	r := mux.NewRouter()
	r.HandleFunc("/", Test)

	server.server = &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return &server, nil
}

func (s *Server) Run() {
	log.Fatal(s.server.ListenAndServe())
}

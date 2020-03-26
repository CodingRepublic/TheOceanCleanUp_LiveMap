package server

import (
	"log"
)

type Config struct {
	Database *DatabaseConfig `yaml:"db"`
	Handler  *HandlerConfig  `yaml:"handler"`
}

type Server struct {
	Database *Database
	Handler  *Handler
}

func Create(config *Config) (*Server, error) {
	server := Server{}
	var err error

	server.Database, err = InitDatabase(config.Database)
	if err != nil {
		return nil, err
	}

	server.Handler, err = InitHandler(config.Handler)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *Server) Run() {
	log.Fatal(s.Handler.server.ListenAndServe())
}

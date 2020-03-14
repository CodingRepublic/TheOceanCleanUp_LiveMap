package main

import (
	_ "github.com/lib/pq"
)

type Config struct {
	Database *DatabaseConfig `yaml:"database"`
}

type Server struct {
	Database *Database
}

func CreateServer(config *Config) (*Server, error) {
	var err error
	server := Server{}

	server.Database, err = GetDatabase(config.Database)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *Server) Run() error {
	return nil
}

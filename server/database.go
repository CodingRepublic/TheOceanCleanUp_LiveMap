package main

import (
	_ "github.com/lib/pq"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type DatabaseConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
}

type Database struct {
	config *DatabaseConfig

	DB *r.Session
}

func InitDatabase(config *DatabaseConfig) (*Database, error) {
	var database Database
	var err error

	database.DB, err = r.Connect(r.ConnectOpts{
		Address: config.Host,
	})
	if err != nil {
		return nil, err
	}

	database.config = config
	return &database, nil
}

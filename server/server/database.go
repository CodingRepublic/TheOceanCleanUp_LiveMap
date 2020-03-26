package server

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

	db *r.Session
}

func InitDatabase(config *DatabaseConfig) (*Database, error) {
	var database Database
	var err error

	database.db, err = r.Connect(r.ConnectOpts{
		Address: config.Host,
	})
	if err != nil {
		return nil, err
	}

	database.config = config
	return &database, nil
}

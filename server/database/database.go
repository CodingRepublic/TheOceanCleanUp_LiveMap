package database

import (
	_ "github.com/lib/pq"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Config struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
}

type Database struct {
	config *Config

	DB *r.Session
}

func Init(config *Config) (*Database, error) {
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

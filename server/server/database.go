package server

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var myDatabase Database

type DatabaseConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
}

type Database struct {
	config *DatabaseConfig

	session *r.Session
}

func InitDatabase(config *DatabaseConfig) error {
	var err error

	myDatabase.session, err = r.Connect(r.ConnectOpts{
		Address: config.Host,
	})
	if err != nil {
		return err
	}

	myDatabase.config = config
	return nil
}

////////////////////////
// DATABASE FUNCTIONS //
////////////////////////

func (db *Database) NewFleetUnit(ID string) error {
	newFleetUnit := FleetUnit{ID, Position{0, 0}}
	_, err := r.DB("theoceancleanup").Table("real_time_fleet").Insert(newFleetUnit).RunWrite(db.session)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateFleetUnit(updatedFleetUnit FleetUnit) error {
	_, err := r.DB("theoceancleanup").Table("real_time_fleet").Update(updatedFleetUnit).RunWrite(db.session)
	if err != nil {
		return err
	}
	return nil
}

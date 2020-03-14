package main

import "database/sql"

type DatabaseConfig struct {
	User  string `yaml:"user"`
	Table string `yaml:"table"`
	Host  string `yaml:"host"`
}

type Database struct {
	DB     *sql.DB
	config *DatabaseConfig
}

func GetDatabase(config *DatabaseConfig) (*Database, error) {
	var database Database
	var err error

	database.DB, err = sql.Open("postgres",
		"postgresql://"+
			config.User+"@"+
			config.Host+
			"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// TEMP
	err = database.ResetDatabase()
	if err != nil {
		return nil, err
	}

	database.config = config
	return &database, nil
}

func (db *Database) ResetDatabase() error {
	var err error

	// Remove tables
	if _, err = db.DB.Exec(
		"DROP DATABASE IF EXISTS postgres"); err != nil {
		return err
	}
	if _, err = db.DB.Exec(
		"DROP DATABASE IF EXISTS defaultdb"); err != nil {
		return err
	}
	if _, err = db.DB.Exec(
		"DROP DATABASE IF EXISTS oceancleanup"); err != nil {
		return err
	}

	// Add tables
	if _, err = db.DB.Exec(
		"CREATE DATABASE IF NOT EXISTS oceancleanup"); err != nil {
		return err
	}

	return nil
}

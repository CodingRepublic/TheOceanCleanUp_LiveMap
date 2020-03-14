package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	LocalEnv      = "local"
	StagingEnv    = "staging"
	ProductionEnv = "production"
)

func getConfig() (*Config, error) {
	var configFilePath string
	var serverConfig Config

	env := os.Getenv("APP_ENV")

	if env == "" {
		return nil, errors.New("APP_ENV not set")
	} else if env == LocalEnv {
		configFilePath = "./config/local/local.config.yaml"
	} else if env == StagingEnv {
		configFilePath = "./config/staging/staging.config.yaml"
	} else if env == ProductionEnv {
		configFilePath = "./config/production/production.config.yaml"
	}

	yamlConfig, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlConfig, &serverConfig)
	if err != nil {
		return nil, err
	}

	return &serverConfig, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Printf("getConfig error: %v\n", err)
		return
	}
	server, err := CreateServer(config)
	if err != nil {
		log.Printf("CreateServer error: %v\n", err)
		return
	}
	err = server.Run()
	if err != nil {
		log.Printf("Run error: %v\n", err)
		return
	}
}

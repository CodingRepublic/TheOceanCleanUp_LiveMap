package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"server/server"
)

const (
	LocalEnv      = "local"
	StagingEnv    = "staging"
	ProductionEnv = "production"
)

func getConfig() (*server.Config, error) {
	var configFilePath string
	var serverConfig server.Config

	env := os.Getenv("APP_ENV")

	if env == "" {
		return nil, errors.New("APP_ENV not set")
	} else if env == LocalEnv {
		configFilePath = "./config/dev/dev.config.yaml"
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

	err = server.Init(config)
	if err != nil {
		log.Printf("Init error: %v\n", err)
		return
	}

	server.Start()
}

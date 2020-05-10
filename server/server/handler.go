package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var myHandler Handler

type HandlerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Handler struct {
	config *HandlerConfig

	server *http.Server
}

func InitHandler(config *HandlerConfig) error {
	r := mux.NewRouter()

	setHandlersApiClient(r)
	setHandlersApiFleet(r)

	myHandler.server = &http.Server{
		Handler: r,
		Addr:    config.Host + ":" + config.Port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	myHandler.config = config
	return nil
}

func (handler *Handler) Start() {
	log.Fatal(handler.server.ListenAndServe())
}

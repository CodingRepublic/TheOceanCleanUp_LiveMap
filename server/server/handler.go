package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HandlerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Handler struct {
	config *HandlerConfig

	server *http.Server
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func InitHandler(config *HandlerConfig) (*Handler, error) {
	var handler Handler

	r := mux.NewRouter()
	r.HandleFunc("/", Home)

	handler.server = &http.Server{
		Handler: r,
		Addr:    config.Host + ":" + config.Port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	handler.config = config
	return &handler, nil
}

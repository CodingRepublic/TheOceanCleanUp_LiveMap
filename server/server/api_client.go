package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func setHandlersApiClient(r *mux.Router) {
	r.HandleFunc("/", Home)
}

package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func _UpdateFleetUnit_(w http.ResponseWriter, r *http.Request) {
	var newFleetUnit FleetUnit

	// TODO json validator
	err := ReadBody(r.Body, &newFleetUnit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = myDatabase.UpdateFleetUnit(newFleetUnit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func setHandlersApiFleet(r *mux.Router) {
	r.HandleFunc("/api/update_fleet_unit", _UpdateFleetUnit_)
}

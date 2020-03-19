package main

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

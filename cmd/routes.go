package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *app) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/location/{orderId}/now", app.getLocations).Methods("POST")
	r.HandleFunc("/location/{orderId}", app.getLocations).Methods("GET") // add ?max
	r.HandleFunc("/location/{orderId}", app.deleteLocation)

	return r 
 }

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *app) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/location/{orderId}/now", app.createLocation).Methods("POST")
	r.HandleFunc("/location/{orderId}", app.getLocations).Methods("GET") 
	r.HandleFunc("/location/{orderId}", app.deleteLocation).Methods("DELETE")

	return r 
 }

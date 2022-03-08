package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zeurd/location/pkg/models"
)

func(app *app) getLocations(w http.ResponseWriter, r *http.Request) { 

	params := mux.Vars(r)
	orderId := params["orderId"]

	var lh models.LocationHistory

	maxQuery := r.URL.Query().Get("max")
	if maxQuery == "" {
		app.locationMapper.GetLocationHistory(orderId)
	} else {
		max, err := strconv.Atoi(maxQuery)
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
		}
		app.locationMapper.GetLocationHistoryLimit(orderId, max)
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lh)
}

func(app *app) createLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	orderId := params["orderId"]

	var location models.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	if err := app.locationMapper.Insert(orderId, location); err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}


func(app *app) deleteLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId := params["orderId"]
	
	if err := app.locationMapper.Delete(orderId); err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *app) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack()) 
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) 
}

func (app *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status) 
}
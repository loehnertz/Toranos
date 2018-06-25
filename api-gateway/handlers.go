package main

import (
	"context"
	"encoding/json"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/micro/go-log"
	"net/http"
)

func getAuthToken(w http.ResponseWriter, r *http.Request) {

}

func registerNewUser(w http.ResponseWriter, r *http.Request) {

}

func availableVehicles(w http.ResponseWriter, r *http.Request) {
	resAvailableVehicles, errAvailableVehicles := fleetMonitor.AvailableVehicles(context.TODO(), &fleet_monitor.Empty{})
	if errAvailableVehicles != nil {
		log.Log(errAvailableVehicles)
		w.Write([]byte(commons.UnknownError.Error()))
	}

	jsonBytes, marshalError := json.Marshal(resAvailableVehicles)
	if marshalError != nil {
		log.Log(marshalError)
		w.Write([]byte(commons.UnknownError.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func createBooking(w http.ResponseWriter, r *http.Request) {

}

func deleteBooking(w http.ResponseWriter, r *http.Request) {

}

func beginRide(w http.ResponseWriter, r *http.Request) {

}

func endRide(w http.ResponseWriter, r *http.Request) {

}

func retrieveStatistics(w http.ResponseWriter, r *http.Request)  {

}

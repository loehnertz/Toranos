package main

import (
	"context"
	"encoding/json"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"net/http"
)

func registerNewUser(w http.ResponseWriter, r *http.Request) {

}

func getAuthToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body getAuthTokenRequest
	decodeError := decoder.Decode(&body)
	if decodeError != nil {
		log.Log(decodeError)
		w.Write([]byte(commons.UnknownError.Error()))
	} else {
		resIssueToken, errIssueToken := userManagement.IssueUserToken(context.TODO(), &user_management.IssueUserTokenRequest{
			Email:    body.Email,
			Password: body.Password,
		})

		if errIssueToken != nil {
			log.Log(errIssueToken)
			w.Write([]byte(commons.UnknownError.Error()))
		} else {
			jsonBytes, marshalError := json.Marshal(resIssueToken)
			if marshalError != nil {
				log.Log(marshalError)
				w.Write([]byte(commons.UnknownError.Error()))
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonBytes)
		}
	}
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

func retrieveStatistics(w http.ResponseWriter, r *http.Request) {

}

package main

import (
	"context"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"net/http"
)

func registerNewUser(w http.ResponseWriter, r *http.Request) {

}

func getAuthToken(w http.ResponseWriter, r *http.Request) {
	deserializedBody, deserializeError := deserialize(new(getAuthTokenRequest), r.Body)
	body := deserializedBody.(*getAuthTokenRequest)

	if deserializeError != nil {
		w.Write([]byte(deserializeError.Error()))
	} else {
		resIssueToken, errIssueToken := userManagement.IssueUserToken(context.TODO(), &user_management.IssueUserTokenRequest{
			Email:    body.Email,
			Password: body.Password,
		})

		if errIssueToken != nil {
			log.Log(errIssueToken)
			w.Write([]byte(commons.UnknownError.Error()))
		} else {
			respondWithJson(&w, resIssueToken)
		}
	}
}

func availableVehicles(w http.ResponseWriter, r *http.Request) {
	resAvailableVehicles, errAvailableVehicles := fleetMonitor.AvailableVehicles(context.TODO(), &fleet_monitor.Empty{})
	if errAvailableVehicles != nil {
		log.Log(errAvailableVehicles)
		w.Write([]byte(commons.UnknownError.Error()))
	}

	respondWithJson(&w, resAvailableVehicles)
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	deserializedBody, _ := deserialize(new(createBookingRequest), r.Body)
	body := deserializedBody.(*createBookingRequest)

	resCreateBooking, errCreateBooking := fleetController.Book(context.TODO(), &fleet_controller.BookingRequest{
		VehicleId:  body.VehicleId,
		CustomerId: "",
	})

	if errCreateBooking != nil {
		w.Write([]byte(errCreateBooking.Error()))
	} else {
		respondWithJson(&w, resCreateBooking)
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {

}

func beginRide(w http.ResponseWriter, r *http.Request) {

}

func endRide(w http.ResponseWriter, r *http.Request) {

}

func retrieveStatistics(w http.ResponseWriter, r *http.Request) {

}

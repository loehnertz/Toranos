package main

import (
	"context"
	gorillacontext "github.com/gorilla/context"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"net/http"
)

func registerNewUser(w http.ResponseWriter, r *http.Request) {
	deserializedBody, deserializeError := deserialize(new(user_management.RegisterCustomerRequest), r.Body)
	body := deserializedBody.(*user_management.RegisterCustomerRequest)

	if deserializeError != nil {
		w.Write([]byte(deserializeError.Error()))
	} else {
		resRegisterCustomer, errRegisterCustomer := userManagement.RegisterCustomer(context.TODO(), &user_management.RegisterCustomerRequest{
			Email:     body.Email,
			Password:  body.Password,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			LicenseId: body.LicenseId,
		})

		if errRegisterCustomer != nil {
			log.Log(errRegisterCustomer)
			w.Write([]byte(commons.UnknownError.Error()))
		} else {
			respondWithJson(&w, resRegisterCustomer)
		}
	}
}

func getAuthToken(w http.ResponseWriter, r *http.Request) {
	deserializedBody, deserializeError := deserialize(new(user_management.IssueUserTokenRequest), r.Body)
	body := deserializedBody.(*user_management.IssueUserTokenRequest)

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
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	deserializedBody, _ := deserialize(new(fleet_controller.BookingRequest), r.Body)
	body := deserializedBody.(*fleet_controller.BookingRequest)

	resCreateBooking, errCreateBooking := fleetController.Book(context.TODO(), &fleet_controller.BookingRequest{
		VehicleId:  body.VehicleId,
		CustomerId: user.Email,
	})

	if errCreateBooking != nil {
		log.Log(errCreateBooking)
		w.Write([]byte(commons.UnknownError.Error()))
	} else {
		respondWithJson(&w, resCreateBooking)
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	deserializedBody, _ := deserialize(new(fleet_controller.UnbookingRequest), r.Body)
	body := deserializedBody.(*fleet_controller.UnbookingRequest)

	resDeleteBooking, errDeleteBooking := fleetController.Unbook(context.TODO(), &fleet_controller.UnbookingRequest{
		VehicleId:  body.VehicleId,
		CustomerId: user.Email,
	})

	if errDeleteBooking != nil {
		log.Log(errDeleteBooking)
		w.Write([]byte(commons.UnknownError.Error()))
	} else {
		respondWithJson(&w, resDeleteBooking)
	}
}

func beginRide(w http.ResponseWriter, r *http.Request) {

}

func endRide(w http.ResponseWriter, r *http.Request) {

}

func retrieveStatistics(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this!
}

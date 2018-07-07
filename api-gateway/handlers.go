package main

import (
	"context"
	gorillacontext "github.com/gorilla/context"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/statistics/proto"
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
		resRegisterCustomer, errRegisterCustomer := userManagementService.RegisterCustomer(context.TODO(), &user_management.RegisterCustomerRequest{
			Email:     body.Email,
			Password:  body.Password,
			FirstName: body.FirstName,
			LastName:  body.LastName,
			LicenseId: body.LicenseId,
		})

		if errRegisterCustomer != nil {
			log.Log(errRegisterCustomer)
			w.Write([]byte(common.UnknownError.Error()))
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
		resIssueToken, errIssueToken := userManagementService.IssueUserToken(context.TODO(), &user_management.IssueUserTokenRequest{
			Email:    body.Email,
			Password: body.Password,
		})

		if errIssueToken != nil {
			log.Log(errIssueToken)
			w.Write([]byte(common.UnknownError.Error()))
		} else {
			respondWithJson(&w, resIssueToken)
		}
	}
}

func availableVehicles(w http.ResponseWriter, r *http.Request) {
	resAvailableVehicles, errAvailableVehicles := fleetMonitorService.AvailableVehicles(context.TODO(), &fleet_monitor.Empty{})
	if errAvailableVehicles != nil {
		log.Log(errAvailableVehicles)
		w.Write([]byte(common.UnknownError.Error()))
	}

	respondWithJson(&w, resAvailableVehicles)
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	deserializedBody, _ := deserialize(new(fleet_controller.BookingRequest), r.Body)
	body := deserializedBody.(*fleet_controller.BookingRequest)

	resCreateBooking, errCreateBooking := fleetControllerService.Book(context.TODO(), &fleet_controller.BookingRequest{
		VehicleId:  body.VehicleId,
		CustomerId: user.Email,
	})

	if errCreateBooking != nil {
		log.Log(errCreateBooking)
		w.Write([]byte(common.UnknownError.Error()))
	} else {
		respondWithJson(&w, resCreateBooking)
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	deserializedBody, _ := deserialize(new(fleet_controller.UnbookingRequest), r.Body)
	body := deserializedBody.(*fleet_controller.UnbookingRequest)

	resDeleteBooking, errDeleteBooking := fleetControllerService.Unbook(context.TODO(), &fleet_controller.UnbookingRequest{
		VehicleId:  body.VehicleId,
		CustomerId: user.Email,
	})

	if errDeleteBooking != nil {
		log.Log(errDeleteBooking)
		w.Write([]byte(common.UnknownError.Error()))
	} else {
		respondWithJson(&w, resDeleteBooking)
	}
}

func beginRide(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	resBeginRide, errBeginRide := fleetControllerService.BeginRide(context.TODO(), &fleet_controller.BeginRideRequest{
		CustomerId: user.Email,
	})

	if errBeginRide != nil {
		log.Log(errBeginRide)
		w.Write([]byte(common.UnknownError.Error()))
	} else {
		respondWithJson(&w, resBeginRide)
	}
}

func endRide(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	resEndRide, errEndRide := fleetControllerService.EndRide(context.TODO(), &fleet_controller.EndRideRequest{
		CustomerId: user.Email,
	})

	if errEndRide != nil {
		log.Log(errEndRide)
		w.Write([]byte(common.UnknownError.Error()))
	} else {
		respondWithJson(&w, resEndRide)
	}
}

func retrieveStatistics(w http.ResponseWriter, r *http.Request) {
	user := gorillacontext.Get(r, "user").(*user_management.AuthenticateUserResponse)

	resBookings, errBookings := statisticsService.RetrieveBookings(context.TODO(), &statistics.RetrieveBookingsRequest{
		UserId: user.Email,
	})

	if errBookings != nil {
		log.Log(errBookings)
		w.Write([]byte(common.UnknownError.Error()))
	} else {
		respondWithJson(&w, resBookings)
	}
}

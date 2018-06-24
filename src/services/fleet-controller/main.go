package main

import (
	"context"
	"fmt"

	"database/sql"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/micro/go-micro"
	"time"
)

const ReservationTimeInSeconds = 900

const DataSource = "user=jloehnertz dbname=toranos_fleet sslmode=disable"

var service micro.Service
var database *sql.DB

type FleetController struct{}

func (fc *FleetController) Book(ctx context.Context, req *fleet_controller.BookingRequest, res *fleet_controller.BookingResponse) error {
	bookingSuccessful, bookingError := book(database, req.VehicleId, req.CustomerId)

	if bookingSuccessful == true {
		res.Successful = true
		res.ReservedTime = ReservationTimeInSeconds
	} else {
		res.Successful = false
		res.Error = bookingError.Error()
	}
	return nil
}

func (fc *FleetController) Unbook(ctx context.Context, req *fleet_controller.UnbookingRequest, res *fleet_controller.UnbookingResponse) error {
	unbookingSuccessful, unbookingError := unbook(database, req.VehicleId, req.CustomerId)

	if unbookingSuccessful == true {
		res.Successful = true
	} else {
		res.Successful = false
		res.Error = unbookingError.Error()
	}

	return nil
}

func (fc *FleetController) RetrieveReservations(ctx context.Context, req *fleet_controller.Empty, res *fleet_controller.RetrieveReservationsResponse) error {
	reservations, reservationsError := retrieveReservations(database)

	if reservationsError == nil {
		res.Reservations = reservations
	} else {
		res.Error = reservationsError.Error()
	}

	return nil
}

func main() {
	var databaseError error
	database, databaseError = sql.Open("postgres", DataSource)
	if databaseError != nil {
		panic(databaseError)
	}

	// Create the service
	service = micro.NewService(
		micro.Name("fleet-controller"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	fleet_controller.RegisterFleetControllerHandler(service.Server(), new(FleetController))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
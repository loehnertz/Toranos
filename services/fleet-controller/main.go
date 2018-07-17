package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/vehicle-gateway/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"time"
)

const DatabaseDriver = "postgres"

var conf config.Config
var database *sql.DB
var service micro.Service
var vehicleGateway vehicle_gateway.VehicleGatewayService

type FleetController struct{}

func (fc *FleetController) Book(ctx context.Context, req *fleet_controller.BookingRequest, res *fleet_controller.BookingResponse) error {
	bookingSuccessful, bookingError := book(req.VehicleId, req.CustomerId)

	if bookingSuccessful == true {
		res.Successful = true
		res.ReservedTime = uint32(common.GetConfigDurationByPath(conf, "service-settings", "fleet-controller", "reservationTime").Seconds())
	} else {
		res.Successful = false
		res.Error = bookingError.Error()
	}

	return nil
}

func (fc *FleetController) Unbook(ctx context.Context, req *fleet_controller.UnbookingRequest, res *fleet_controller.UnbookingResponse) error {
	unbookingSuccessful, unbookingError := unbook(req.VehicleId, req.CustomerId)

	if unbookingSuccessful == true {
		res.Successful = true
	} else {
		res.Successful = false
		res.Error = unbookingError.Error()
	}

	return nil
}

func (fc *FleetController) BeginRide(ctx context.Context, req *fleet_controller.BeginRideRequest, res *fleet_controller.BeginRideResponse) error {
	beginRideSuccessful, beginRideError := beginRide(req.CustomerId)

	if beginRideSuccessful == true {
		res.Successful = true
	} else {
		res.Successful = false
		res.Error = beginRideError.Error()
	}

	return nil
}

func (fc *FleetController) EndRide(ctx context.Context, req *fleet_controller.EndRideRequest, res *fleet_controller.EndRideResponse) error {
	endRideSuccessful, endRideError := endRide(req.CustomerId)

	if endRideSuccessful == true {
		res.Successful = true
	} else {
		res.Successful = false
		res.Error = endRideError.Error()
	}

	return nil
}

func (fc *FleetController) RetrieveReservations(ctx context.Context, req *fleet_controller.Empty, res *fleet_controller.RetrieveReservationsResponse) error {
	reservations, reservationsError := retrieveReservations()

	if reservationsError == nil {
		res.Reservations = reservations
	} else {
		res.Error = reservationsError.Error()
	}

	return nil
}

func (fc *FleetController) RetrieveUnbilledBookings(ctx context.Context, req *fleet_controller.Empty, res *fleet_controller.RetrieveUnbilledBookingsResponse) error {
	bookings, bookingsError := retrieveUnbilledBookings()

	if bookingsError == nil {
		res.Bookings = bookings
	}

	return bookingsError
}

func (fc *FleetController) RetrieveBilledBookingsOfCustomer(ctx context.Context, req *fleet_controller.RetrieveBilledBookingsOfCustomerRequest, res *fleet_controller.RetrieveBilledBookingsOfCustomerResponse) error {
	bookings, bookingsError := retrieveBilledBookingsOfCustomer(req.UserId)

	if bookingsError == nil {
		res.Bookings = bookings
	}

	return bookingsError
}

func (fc *FleetController) AddInvoiceToBooking(ctx context.Context, req *fleet_controller.AddInvoiceToBookingRequest, res *fleet_controller.AddInvoiceToBookingResponse) error {
	addInvoiceSuccessful, addInvoiceError := addInvoiceToBooking(req.BookingId, req.InvoiceId)

	if addInvoiceSuccessful {
		res.Successful = true
	}

	return addInvoiceError
}

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Connect the database
	var databaseError error
	database, databaseError = sql.Open(
		DatabaseDriver,
		common.ConstructPostgresDataSourceString(
			common.GetConfigStringByPath(conf, "databases", "postgres", "fleet", "name"),
			common.GetConfigStringByPath(conf, "databases", "postgres", "fleet", "user"),
			common.GetConfigStringByPath(conf, "databases", "postgres", "fleet", "ssl"),
		),
	)
	if databaseError != nil {
		panic(databaseError)
	}

	// Create the service
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "fleet-controller")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	serviceClient := service.Client()
	serviceClient.Init(client.Retries(3))

	// Initialize the service clients
	vehicleGateway = vehicle_gateway.NewVehicleGatewayService(
		common.GetConfigStringByPath(conf, "service-names", "vehicle-gateway"),
		service.Client(),
	)

	// Register the handler
	fleet_controller.RegisterFleetControllerHandler(service.Server(), new(FleetController))

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

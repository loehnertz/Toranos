package main

import (
	"context"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/statistics/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"time"
)

var conf config.Config
var service micro.Service
var fleetControllerService fleet_controller.FleetControllerService

type Statistics struct{}

func (st *Statistics) RetrieveBookings(ctx context.Context, req *statistics.RetrieveBookingsRequest, res *statistics.RetrieveBookingsResponse) error {
	bookings, bookingsError := retrieveAllBookingsOfCustomer(req.UserId)

	if bookingsError == nil {
		res.Bookings = bookings
	}

	return bookingsError
}

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Create the service
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "statistics")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	statistics.RegisterStatisticsHandler(service.Server(), new(Statistics))

	// Initialize the service clients
	fleetControllerService = fleet_controller.NewFleetControllerService(
		common.GetConfigStringByPath(conf, "service-names", "fleet-controller"),
		service.Client(),
	)

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

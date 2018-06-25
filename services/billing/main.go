package main

import (
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	"time"
)

var service micro.Service
var fleetController fleet_controller.FleetControllerService

func main() {
	// Create the service
	service = micro.NewService(
		micro.Name(config.BookingName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Initialize the service clients
	fleetController = fleet_controller.NewFleetControllerService(config.FleetControllerName, service.Client())

	// Initialize all the tasks
	scheduler := cron.New()

	scheduler.AddFunc(config.CheckForBookingsToBillInterval, checkForBookingsToBill)

	// Start all the tasks
	scheduler.Start()

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

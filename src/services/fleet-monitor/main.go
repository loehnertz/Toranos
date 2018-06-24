package main

import (
	"context"
	"fmt"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/src/services/fleet-monitor/proto"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	"time"
)

var service micro.Service
var fleetController fleet_controller.FleetControllerService

type FleetMonitor struct{}

func (fm *FleetMonitor) AvailableVehicles(ctx context.Context, req *fleet_monitor.Empty, res *fleet_monitor.AvailableVehiclesResponse) error {
	res.Vehicles = retrieveAvailableVehicles()

	return nil
}

func main() {
	// Create the service
	service = micro.NewService(
		micro.Name(config.FleetMonitorName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	fleet_monitor.RegisterFleetMonitorHandler(service.Server(), new(FleetMonitor))

	// Initialize the FleetController client
	fleetController = fleet_controller.NewFleetControllerService(config.FleetControllerName, service.Client())

	// Initialize all the tasks
	scheduler := cron.New()

	scheduler.AddFunc(config.CheckForExpiredReservationsInterval, checkForExpiredReservations)

	// Start all the tasks
	scheduler.Start()

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

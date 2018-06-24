package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/micro/go-micro"
)

var fleetController fleet_controller.FleetControllerService

func main() {
	service := micro.NewService(micro.Name("fleet-monitor"))
	service.Init()

	fleetController = fleet_controller.NewFleetControllerService("fleet-controller", service.Client())

	gocron.Every(config.CheckForExpiredReservationsIntervalInSeconds).Seconds().Do(checkForExpiredReservations)
	<-gocron.Start()
}

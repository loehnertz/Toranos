package main

import (
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	"time"
)

var conf config.Config
var service micro.Service
var fleetController fleet_controller.FleetControllerService

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Create the service
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "billing")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Initialize the service clients
	fleetController = fleet_controller.NewFleetControllerService(
		common.GetConfigStringByPath(conf, "service-names", "fleet-controller"),
		service.Client(),
	)

	// Initialize all the tasks
	scheduler := cron.New()

	scheduler.AddFunc(common.GetConfigStringByPath(conf, "crons", "checkForBookingsToBill"), checkForBookingsToBill)

	// Start all the tasks
	scheduler.Start()

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

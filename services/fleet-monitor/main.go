package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/telemetry/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	"time"
)

var conf config.Config
var redisClient *redis.Client
var service micro.Service
var fleetController fleet_controller.FleetControllerService
var telemetryService telemetry.TelemetryService

type FleetMonitor struct{}

func (fm *FleetMonitor) AvailableVehicles(ctx context.Context, req *fleet_monitor.Empty, res *fleet_monitor.AvailableVehiclesResponse) error {
	res.Vehicles = retrieveAvailableVehicles()

	return nil
}

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Initialize a Redis client
	redisClient = common.InitRedisClient(
		common.GetConfigStringByPath(conf, "caching", "redis", "host"),
		common.GetConfigStringByPath(conf, "caching", "redis", "password"),
		common.GetConfigIntByPath(conf, "caching", "redis", "databaseId"),
	)

	// Create the service
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "fleet-monitor")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	fleet_monitor.RegisterFleetMonitorHandler(
		service.Server(),
		new(FleetMonitor),
	)

	// Initialize the service clients
	fleetController = fleet_controller.NewFleetControllerService(
		common.GetConfigStringByPath(conf, "service-names", "fleet-controller"),
		service.Client(),
	)
	telemetryService = telemetry.NewTelemetryService(
		common.GetConfigStringByPath(conf, "service-names", "telemetry"),
		service.Client(),
	)

	// Initialize all the tasks
	scheduler := cron.New()

	scheduler.AddFunc(common.GetConfigStringByPath(conf, "crons", "checkForExpiredReservations"), checkForExpiredReservations)

	// Start all the tasks
	scheduler.Start()

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

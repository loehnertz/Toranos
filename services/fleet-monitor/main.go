package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/telemetry/proto"
	"github.com/micro/go-micro"
	"github.com/robfig/cron"
	"time"
)

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
	// Initialize a Redis client
	redisClient = commons.InitRedisClient(commons.RedisHostAddress, "", commons.RedisDatabaseId)

	// Create the service
	service = micro.NewService(
		micro.Name(config.FleetMonitorName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	fleet_monitor.RegisterFleetMonitorHandler(service.Server(), new(FleetMonitor))

	// Initialize the service clients
	fleetController = fleet_controller.NewFleetControllerService(config.FleetControllerName, service.Client())
	telemetryService = telemetry.NewTelemetryService(config.TelemetryName, service.Client())

	// Initialize all the tasks
	scheduler := cron.New()

	scheduler.AddFunc(config.CheckForExpiredReservationsInterval, checkForExpiredReservations)

	// Start all the tasks
	scheduler.Start()

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

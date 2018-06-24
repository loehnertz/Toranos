package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/telemetry/proto"
	"github.com/micro/go-micro"
	"time"
)

const RedisAllVehiclesKey = "all_vehicles"

var redisClient *redis.Client
var service micro.Service

type Telemetry struct{}

func (tm *Telemetry) AllVehicles(ctx context.Context, req *telemetry.Empty, res *telemetry.AllVehiclesResponse) error {
	res.Vehicles = retrieveAllVehicles(redisClient)

	return nil
}

func main() {
	redisClient = commons.InitRedisClient(commons.RedisHostAddress, "", commons.RedisDatabaseId)

	// Create the service
	service = micro.NewService(
		micro.Name(config.TelemetryName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	telemetry.RegisterTelemetryHandler(service.Server(), new(Telemetry))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

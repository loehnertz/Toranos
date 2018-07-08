package main

import (
	"context"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/vehicle-gateway/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"time"
)

var conf config.Config
var service micro.Service

type VehicleGateway struct{}

func (vg *VehicleGateway) ReachVehicle(ctx context.Context, req *vehicle_gateway.ReachVehicleRequest, res *vehicle_gateway.ReachVehicleResponse) error {
	res.Successful = reachVehicle(req.VehicleId, req.Method)

	return nil
}

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Create the service client
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "vehicle-gateway")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	vehicle_gateway.RegisterVehicleGatewayHandler(service.Server(), new(VehicleGateway))

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}

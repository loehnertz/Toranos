#!/bin/sh

rm -rf ./bin/
mkdir ./bin/

go build -o ./bin/api_gateway ./api-gateway
go build -o ./bin/billing ./services/billing
go build -o ./bin/fleet_controller ./services/fleet-controller
go build -o ./bin/fleet_monitor ./services/fleet-monitor
go build -o ./bin/statistics ./services/statistics
go build -o ./bin/telemetry ./services/telemetry
go build -o ./bin/user_management ./services/user-management
go build -o ./bin/vehicle_gateway ./vehicle-gateway

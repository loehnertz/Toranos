#!/bin/sh

rm -rf ./logs/
mkdir ./logs/

nohup consul agent -dev -ui -bind=127.0.0.1 -data-dir=/tmp/consul > ./logs/consul.out &
sleep 4
cat ./config.consul.json | consul kv import - > ./logs/consul.out
sleep 1

nohup ./bin/api_gateway > ./logs/api_gateway.out &
nohup ./bin/billing > ./logs/billing.out &
nohup ./bin/fleet_controller > ./logs/fleet_controller.out &
nohup ./bin/fleet_monitor > ./logs/fleet_monitor.out &
nohup ./bin/statistics > ./logs/statistics.out &
nohup ./bin/telemetry > ./logs/telemetry.out &
nohup ./bin/user_management > ./logs/user_management.out &
nohup ./bin/vehicle_gateway > ./logs/vehicle_gateway.out &

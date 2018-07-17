#!/bin/sh

rm -rf ./logs/
mkdir ./logs/

# Consul is run in `-dev` mode. This should not be the case in a production environment
nohup consul agent -dev -ui -bind=127.0.0.1 -data-dir=/tmp/consul > ./logs/consul.out 2>&1 &
sleep 4
cat ./config.consul.json | consul kv import - > ./logs/consul.out
sleep 1

nohup ./bin/api_gateway > ./logs/api_gateway.out 2>&1 &
nohup ./bin/billing > ./logs/billing.out 2>&1 &
nohup ./bin/fleet_controller > ./logs/fleet_controller.out 2>&1 &
nohup ./bin/fleet_monitor > ./logs/fleet_monitor.out 2>&1 &
nohup ./bin/statistics > ./logs/statistics.out 2>&1 &
nohup ./bin/telemetry > ./logs/telemetry.out 2>&1 &
nohup ./bin/user_management > ./logs/user_management.out 2>&1 &
nohup ./bin/vehicle_gateway > ./logs/vehicle_gateway.out 2>&1 &

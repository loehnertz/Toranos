package main

import (
	"context"
	"encoding/json"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/src/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/src/services/telemetry/proto"
	"github.com/micro/go-log"
	"time"
)

const RedisAvailableVehiclesKey = "available_vehicles"

func retrieveAvailableVehicles() (vehicles []*fleet_monitor.AvailableVehiclesResponse_Vehicle) {
	result, redisGetError := redisClient.Get(RedisAvailableVehiclesKey).Result()

	if redisGetError == nil {
		var availableVehicles []fleet_monitor.AvailableVehiclesResponse_Vehicle
		jsonUnmarshalError := json.Unmarshal([]byte(result), &availableVehicles)
		if jsonUnmarshalError != nil {
			log.Log(jsonUnmarshalError)
			vehicles = getAvailableVehiclesWhileRefreshingRedisCache()
		}

		for i := range availableVehicles {
			vehicles = append(vehicles, &availableVehicles[i])
		}
	} else {
		vehicles = getAvailableVehiclesWhileRefreshingRedisCache()
	}

	return
}

func getAvailableVehiclesWhileRefreshingRedisCache() (vehicles []*fleet_monitor.AvailableVehiclesResponse_Vehicle) {
	availableVehicles := determineAvailableVehicles()

	writeAvailableVehiclesIntoRedisCache(availableVehicles)

	for i := range availableVehicles {
		vehicles = append(vehicles, &availableVehicles[i])
	}

	return
}

func determineAvailableVehicles() (vehicles []fleet_monitor.AvailableVehiclesResponse_Vehicle) {
	resReservations, errReservations := fleetController.RetrieveReservations(context.TODO(), &fleet_controller.Empty{})
	resAllVehicles, errAllVehicles := telemetryService.AllVehicles(context.TODO(), &telemetry.Empty{})

	if errReservations == nil && errAllVehicles == nil {
		var reservedVehicles []string
		for i := range resReservations.Reservations {
			reservedVehicles = append(reservedVehicles, resReservations.Reservations[i].Vehicle)
		}

		for i := range resAllVehicles.Vehicles {
			vehicle := resAllVehicles.Vehicles[i]

			contains := commons.SliceOfStringsContains(reservedVehicles, vehicle.VehicleId)
			if contains == false {
				vehicles = append(vehicles, fleet_monitor.AvailableVehiclesResponse_Vehicle{
					VehicleId:                      vehicle.VehicleId,
					Location:                       vehicle.Location,
					ApproximateRadialRangeInMeters: calculateApproximateRadialRange(vehicle.Battery),
				})
			}
		}
	} else {
		log.Log(errReservations, errAllVehicles)
	}

	return
}

func writeAvailableVehiclesIntoRedisCache(structure interface{}) {
	redisSetError := redisClient.Set(RedisAvailableVehiclesKey, commons.StringifyIntoJson(structure), config.RedisAvailableVehiclesExpirationTimeInSeconds*time.Second).Err()
	if redisSetError != nil {
		log.Log(redisSetError)
	}
}

func calculateApproximateRadialRange(battery uint32) uint32 {
	return battery // TODO: Implement this!
}

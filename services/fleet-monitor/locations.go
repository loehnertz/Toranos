package main

import (
	"context"
	"encoding/json"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/telemetry/proto"
	"github.com/micro/go-log"
)

func retrieveAvailableVehicles() (vehicles []*fleet_monitor.AvailableVehiclesResponse_Vehicle) {
	result, redisGetError := redisClient.Get(common.GetConfigStringByPath(conf, "caching", "keys", "availableVehicles")).Result()

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

			contains := common.SliceOfStringsContains(reservedVehicles, vehicle.VehicleId)
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
	redisSetError := redisClient.Set(
		common.GetConfigStringByPath(conf, "caching", "keys", "availableVehicles"),
		common.StringifyIntoJson(structure),
		common.GetConfigDurationByPath(conf, "caching", "ttls", "availableVehicles"),
	).Err()
	if redisSetError != nil {
		log.Log(redisSetError)
	}
}

func calculateApproximateRadialRange(battery uint32) uint32 {
	return battery * 100 // TODO: Implement this!
}

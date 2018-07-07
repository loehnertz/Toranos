package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/telemetry/proto"
	"github.com/micro/go-log"
)

func retrieveAllVehicles(redisClient *redis.Client) (vehicles []*telemetry.AllVehiclesResponse_Vehicle) {
	result, redisRetrievalError := redisClient.Get(RedisAllVehiclesKey).Result()

	if redisRetrievalError != nil {
		log.Log(redisRetrievalError)
	} else {
		var allVehicles []common.Vehicle
		jsonUnmarshalError := json.Unmarshal([]byte(result), &allVehicles)
		if jsonUnmarshalError != nil {
			log.Log(jsonUnmarshalError)
		} else {
			for i := range allVehicles {
				vehicle := allVehicles[i]
				vehicles = append(vehicles, &telemetry.AllVehiclesResponse_Vehicle{
					VehicleId: vehicle.Id,
					Location:  vehicle.Location,
					Battery:   vehicle.Battery,
				})
			}
		}
	}

	return
}

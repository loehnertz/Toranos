package main

import (
	"context"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/src/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/src/services/telemetry/proto"
	"github.com/micro/go-log"
)

func retrieveAvailableVehicles() (vehicles []*fleet_monitor.AvailableVehiclesResponse_Vehicle) {
	return determineAvailableVehicles()
}

func determineAvailableVehicles() (vehicles []*fleet_monitor.AvailableVehiclesResponse_Vehicle) {
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
				vehicles = append(vehicles, &fleet_monitor.AvailableVehiclesResponse_Vehicle{
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

func calculateApproximateRadialRange(battery uint32) uint32 {
	return battery // TODO: Implement this!
}

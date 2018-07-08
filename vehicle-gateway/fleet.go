package main

import "fmt"

func reachVehicle(vehicleId string, method string) (successful bool) {
	fmt.Printf("The vehicle with the ID '%v' was reached with the method '%v'\n", vehicleId, method)

	successful = true

	return
}

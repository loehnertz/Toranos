package main

import (
	"context"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/statistics/proto"
	"github.com/micro/go-log"
)

func retrieveAllBookingsOfCustomer(customer string) (bookings []*statistics.RetrieveBookingsResponse_Booking, err error) {
	resBookings, errBookings := fleetControllerService.RetrieveBilledBookingsOfCustomer(context.TODO(), &fleet_controller.RetrieveBilledBookingsOfCustomerRequest{
		UserId: customer,
	})

	if errBookings != nil {
		log.Log(errBookings)
		return
	}

	for i := range resBookings.Bookings {
		booking := resBookings.Bookings[i]

		bookings = append(bookings, &statistics.RetrieveBookingsResponse_Booking{
			CreatedAt:      booking.CreatedAt,
			VehicleId:      booking.VehicleId,
			AmountPaid:     0.0, // TODO: Add a RPC towards the `Billing` service here
			DistanceDriven: booking.DistanceDriven,
			TimeDriven:     booking.TimeDriven,
		})
	}

	return
}

package main

import (
	"context"
	"fmt"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"math/rand"
	"time"
)

func checkForBookingsToBill() {
	resBookings, errBookings := fleetController.RetrieveUnbilledBookings(context.TODO(), &fleet_controller.Empty{})

	if errBookings != nil {
		log.Log(errBookings)
	} else {
		for i := range resBookings.Bookings {
			booking := resBookings.Bookings[i]

			billingSuccessful, invoiceId := createInvoice(booking.Id, booking.DistanceDriven, booking.TimeDriven)

			if !billingSuccessful {
				log.Log("Billing for booking '%v' failed!", booking.Id)
			} else {
				resBilledBooking, errBilledBooking := fleetController.AddInvoiceToBooking(context.TODO(), &fleet_controller.AddInvoiceToBookingRequest{
					BookingId: booking.Id,
					InvoiceId: invoiceId,
				})
				if !resBilledBooking.Successful || errBilledBooking != nil {
					log.Log(errBilledBooking)
				}
			}
		}
		fmt.Printf("Task '%v' finished @ %v \n", "checkForBookingsToBill", time.Now())
	}
}

func createInvoice(bookingId uint32, distanceDriven uint32, timeDriven uint32) (successful bool, invoiceId string) {
	// TODO: Connect third-party billing provider
	invoiceId = "I-" + string(uint32(rand.Intn(100000))+bookingId)
	successful = true
	fmt.Printf("Invoice for booking '%v' was successfully created! \n The user drove %v meters within %v minutes. \n", bookingId, distanceDriven, timeDriven)
	return
}

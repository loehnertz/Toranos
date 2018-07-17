package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/loehnertz/Toranos/common"
	"github.com/loehnertz/Toranos/vehicle-gateway/proto"
	"github.com/micro/go-log"
	"strings"
	"time"
)

const ReservationsOnVehicleId = "SELECT * FROM bookings WHERE vehicle = $1 AND (status = $2 OR status = $3)"
const ReservationsOnCustomerId = "SELECT * FROM bookings WHERE customer = $1 AND (status = $2 OR status = $3)"
const ReservationsOnVehicleIdAndCustomerId = "SELECT id FROM bookings WHERE vehicle = $1 AND customer = $2 AND status = $3"
const InsertNewBooking = "INSERT INTO bookings (created_at, vehicle, customer) VALUES ($1, $2, $3)"
const CancelExistingBooking = "UPDATE bookings SET status = $1 WHERE vehicle = $2 AND customer = $3 AND status = $4"
const UpdateBookingOnCustomerIdWithCertainStatus = "UPDATE bookings SET status = $1 WHERE customer = $2 AND status = $3"
const BookedVehicleOfCustomerIdWithCertainStatus = "SELECT vehicle FROM bookings WHERE customer = $1 AND status = $2"
const UpdateInvoiceOfBooking = "UPDATE bookings SET invoice = $1 WHERE id = $2 AND (status = $3 OR status = $4)"

var vehicleAlreadyBookedError = errors.New("vehicle already booked")
var customerAlreadyBookedError = errors.New("customer already booked a vehicle")
var vehicleWasNotBookedError = errors.New("the selected vehicle was not booked beforehand by the customer")
var bookingCouldNotBeDeletedError = errors.New("booking could not be deleted for an unknown reason")
var beginningRideFailedError = errors.New("beginning the ride failed")
var endingRideFailedError = errors.New("ending the ride failed")

func book(vehicleId string, customerId string) (booked bool, err error) {
	BookingsOnVehicleIdRows := database.QueryRow(
		ReservationsOnVehicleId,
		vehicleId,
		getStatusKeyByName(Reserved),
		getStatusKeyByName(Driving),
	)
	BookingsOnVehicleIdError := BookingsOnVehicleIdRows.Scan()
	BookingsOnCustomerIdRows := database.QueryRow(
		ReservationsOnCustomerId,
		customerId,
		getStatusKeyByName(Reserved),
		getStatusKeyByName(Driving),
	)
	BookingsOnCustomerIdError := BookingsOnCustomerIdRows.Scan()

	if BookingsOnVehicleIdError == sql.ErrNoRows {
		if BookingsOnCustomerIdError == sql.ErrNoRows {
			_, insertError := database.Exec(InsertNewBooking, time.Now().UTC(), vehicleId, customerId)
			if insertError == nil {
				booked = true
			} else {
				err = insertError
				log.Log(insertError)
			}
		} else if strings.Contains(BookingsOnCustomerIdError.Error(), common.FoundRowsErrorSubstring) {
			err = customerAlreadyBookedError
		} else {
			err = BookingsOnCustomerIdError
		}
	} else {
		if strings.Contains(BookingsOnVehicleIdError.Error(), common.FoundRowsErrorSubstring) {
			err = vehicleAlreadyBookedError
		} else {
			err = BookingsOnVehicleIdError
		}
	}

	return
}

func unbook(vehicleId string, customerId string) (unbooked bool, err error) {
	var id int
	row := database.QueryRow(
		ReservationsOnVehicleIdAndCustomerId,
		vehicleId,
		customerId,
		getStatusKeyByName(Reserved),
	)
	selectError := row.Scan(&id)

	if selectError == nil {
		_, deleteError := database.Exec(
			CancelExistingBooking,
			getStatusKeyByName(Canceled),
			vehicleId,
			customerId,
			getStatusKeyByName(Reserved),
		)
		if deleteError != nil {
			log.Log(deleteError)
			err = bookingCouldNotBeDeletedError
		} else {
			unbooked = true
		}
	} else if selectError == sql.ErrNoRows {
		err = vehicleWasNotBookedError
	} else {
		log.Log(selectError)
		err = selectError
	}

	return
}

func beginRide(customerId string) (beginRideSuccessful bool, err error) {
	var vehicle string
	row := database.QueryRow(
		BookedVehicleOfCustomerIdWithCertainStatus,
		customerId,
		getStatusKeyByName(Reserved),
	)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = beginningRideFailedError
	} else {
		_, updateError := database.Exec(
			UpdateBookingOnCustomerIdWithCertainStatus,
			getStatusKeyByName(Driving),
			customerId,
			getStatusKeyByName(Reserved),
		)
		if updateError != nil {
			log.Log(updateError)
			err = beginningRideFailedError
		} else {
			fmt.Printf("Unlocking vehicle '%v' \n", vehicle)

			resReachVehicle, errReachVehicle := vehicleGateway.ReachVehicle(context.TODO(), &vehicle_gateway.ReachVehicleRequest{
				VehicleId: vehicle,
				Method:    "unlock",
			})
			if errReachVehicle != nil {
				log.Log(errReachVehicle)
			}

			beginRideSuccessful = resReachVehicle.Successful
		}
	}

	return
}

func endRide(customerId string) (endRideSuccessful bool, err error) {
	var vehicle string
	row := database.QueryRow(
		BookedVehicleOfCustomerIdWithCertainStatus,
		customerId,
		getStatusKeyByName(Driving),
	)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = endingRideFailedError
	} else {
		_, updateError := database.Exec(
			UpdateBookingOnCustomerIdWithCertainStatus,
			getStatusKeyByName(Done),
			customerId,
			getStatusKeyByName(Driving),
		)
		if updateError != nil {
			log.Log(updateError)
			err = endingRideFailedError
		} else {
			fmt.Printf("Locking vehicle '%v' \n", vehicle)

			resReachVehicle, errReachVehicle := vehicleGateway.ReachVehicle(context.TODO(), &vehicle_gateway.ReachVehicleRequest{
				VehicleId: vehicle,
				Method:    "lock",
			})
			if errReachVehicle != nil {
				log.Log(errReachVehicle)
			}

			endRideSuccessful = resReachVehicle.Successful
		}
	}

	return
}

func addInvoiceToBooking(bookingId uint32, invoiceId string) (successful bool, err error) {
	_, updateError := database.Exec(
		UpdateInvoiceOfBooking,
		invoiceId,
		bookingId,
		getStatusKeyByName(Done),
		getStatusKeyByName(Canceled),
	)
	if updateError != nil {
		log.Log(updateError)
		err = common.UnknownError
	} else {
		successful = true
	}

	return
}

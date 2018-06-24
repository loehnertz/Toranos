package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/loehnertz/toranos/src/config"
	"github.com/micro/go-log"
	"strings"
	"time"
)

const BookingsOnVehicleId = "SELECT * FROM bookings WHERE vehicle = $1"
const BookingsOnCustomerId = "SELECT * FROM bookings WHERE customer = $1"
const BookingsOnVehicleIdAndCustomerId = "SELECT id FROM bookings WHERE vehicle = $1 AND customer = $2"
const InsertNewBooking = "INSERT INTO bookings (created_at, vehicle, customer) VALUES ($1, $2, $3)"
const DeleteBooking = "DELETE FROM bookings WHERE vehicle = $1 AND customer = $2"
const UpdateBookingOnCustomerId = "UPDATE bookings SET status = $1 WHERE customer = $2"
const BookedVehicleOfCustomerIdWithCertainStatus = "SELECT vehicle FROM bookings WHERE customer = $1 AND status = $2"

var vehicleAlreadyBookedError = errors.New("vehicle already booked")
var customerAlreadyBookedError = errors.New("customer already booked a vehicle")
var vehicleWasNotBookedError = errors.New("the selected vehicle was not booked beforehand by the customer")
var bookingCouldNotBeDeletedError = errors.New("booking could not be deleted for an unknown reason")
var beginningRideFailedError = errors.New("beginning the ride failed")
var endingRideFailedError = errors.New("ending the ride failed")

func book(database *sql.DB, vehicleId string, customerId string) (booked bool, err error) {
	BookingsOnVehicleIdRows := database.QueryRow(BookingsOnVehicleId, vehicleId)
	BookingsOnVehicleIdError := BookingsOnVehicleIdRows.Scan()
	BookingsOnCustomerIdRows := database.QueryRow(BookingsOnCustomerId, customerId)
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
		} else if strings.Contains(BookingsOnCustomerIdError.Error(), commons.FoundRowsErrorSubstring) {
			err = customerAlreadyBookedError
		} else {
			err = BookingsOnCustomerIdError
		}
	} else {
		if strings.Contains(BookingsOnVehicleIdError.Error(), commons.FoundRowsErrorSubstring) {
			err = vehicleAlreadyBookedError
		} else {
			err = BookingsOnVehicleIdError
		}
	}

	return
}

func unbook(database *sql.DB, vehicleId string, customerId string) (unbooked bool, err error) {
	var id int
	row := database.QueryRow(BookingsOnVehicleIdAndCustomerId, vehicleId, customerId)
	selectError := row.Scan(&id)

	if selectError == nil {
		_, deleteError := database.Exec(DeleteBooking, vehicleId, customerId)
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

func beginRide(database *sql.DB, customerId string) (beginRideSuccessful bool, err error) {
	var vehicle string
	row := database.QueryRow(BookedVehicleOfCustomerIdWithCertainStatus, customerId, config.StatusReserved)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = beginningRideFailedError
	} else {
		_, updateError := database.Exec(UpdateBookingOnCustomerId, config.StatusDriving, customerId)
		if updateError != nil {
			log.Log(updateError)
			err = beginningRideFailedError
		} else {
			fmt.Printf("Unlocking vehicle '%v' \n", vehicle)
			beginRideSuccessful = true
		}
	}

	return
}

func endRide(database *sql.DB, customerId string) (endRideSuccessful bool, err error) {
	var vehicle string
	row := database.QueryRow(BookedVehicleOfCustomerIdWithCertainStatus, customerId, config.StatusDriving)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = endingRideFailedError
	} else {
		_, updateError := database.Exec(UpdateBookingOnCustomerId, config.StatusDone, customerId)
		if updateError != nil {
			log.Log(updateError)
			err = endingRideFailedError
		} else {
			fmt.Printf("Locking vehicle '%v' \n", vehicle)
			endRideSuccessful = true
		}
	}

	return
}

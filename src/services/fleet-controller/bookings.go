package main

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/micro/go-log"
	"strings"
	"time"
)

const BookingsOnVehicleId = "SELECT * FROM bookings WHERE vehicle = $1"
const BookingsOnCustomerId = "SELECT * FROM bookings WHERE customer = $1"
const BookingsOnVehicleIdAndCustomerId = "SELECT id FROM bookings WHERE vehicle = $1 AND customer = $2"
const InsertNewBooking = "INSERT INTO bookings (created_at, vehicle, customer) VALUES ($1, $2, $3)"
const DeleteBooking = "DELETE FROM bookings WHERE vehicle = $1 AND customer = $2"

var vehicleAlreadyBookedError = errors.New("vehicle already booked")
var customerAlreadyBookedError = errors.New("customer already booked a vehicle")
var vehicleWasNotBookedError = errors.New("the selected vehicle was not booked beforehand by the customer")
var bookingCouldNotBeDeletedError = errors.New("booking could not be deleted for an unknown reason")

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
	rows := database.QueryRow(BookingsOnVehicleIdAndCustomerId, vehicleId, customerId)
	selectError := rows.Scan(&id)

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

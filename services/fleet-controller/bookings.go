package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/config"
	"github.com/micro/go-log"
	"strings"
	"time"
)

const ReservationsOnVehicleId = "SELECT * FROM bookings WHERE vehicle = $1 AND (status = $2 OR status = $3)"
const ReservationsOnCustomerId = "SELECT * FROM bookings WHERE customer = $1 AND (status = $2 OR status = $3)"
const ReservationsOnVehicleIdAndCustomerId = "SELECT id FROM bookings WHERE vehicle = $1 AND customer = $2 AND status = $3"
const InsertNewBooking = "INSERT INTO bookings (created_at, vehicle, customer) VALUES ($1, $2, $3)"
const DeleteExistingBooking = "DELETE FROM bookings WHERE vehicle = $1 AND customer = $2 AND status = $3"
const UpdateBookingOnCustomerIdWithCertainStatus = "UPDATE bookings SET status = $1 WHERE customer = $2 AND status = $3"
const BookedVehicleOfCustomerIdWithCertainStatus = "SELECT vehicle FROM bookings WHERE customer = $1 AND status = $2"
const UpdateInvoiceOfBooking = "UPDATE bookings SET invoice = $1 WHERE id = $2 AND status = $3"

var vehicleAlreadyBookedError = errors.New("vehicle already booked")
var customerAlreadyBookedError = errors.New("customer already booked a vehicle")
var vehicleWasNotBookedError = errors.New("the selected vehicle was not booked beforehand by the customer")
var bookingCouldNotBeDeletedError = errors.New("booking could not be deleted for an unknown reason")
var beginningRideFailedError = errors.New("beginning the ride failed")
var endingRideFailedError = errors.New("ending the ride failed")

func book(vehicleId string, customerId string) (booked bool, err error) {
	BookingsOnVehicleIdRows := database.QueryRow(ReservationsOnVehicleId, vehicleId, config.StatusReserved, config.StatusDriving)
	BookingsOnVehicleIdError := BookingsOnVehicleIdRows.Scan()
	BookingsOnCustomerIdRows := database.QueryRow(ReservationsOnCustomerId, customerId, config.StatusReserved, config.StatusDriving)
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

func unbook(vehicleId string, customerId string) (unbooked bool, err error) {
	var id int
	row := database.QueryRow(ReservationsOnVehicleIdAndCustomerId, vehicleId, customerId, config.StatusReserved)
	selectError := row.Scan(&id)

	if selectError == nil {
		_, deleteError := database.Exec(DeleteExistingBooking, vehicleId, customerId, config.StatusReserved)
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
	row := database.QueryRow(BookedVehicleOfCustomerIdWithCertainStatus, customerId, config.StatusReserved)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = beginningRideFailedError
	} else {
		_, updateError := database.Exec(UpdateBookingOnCustomerIdWithCertainStatus, config.StatusDriving, customerId, config.StatusReserved)
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

func endRide(customerId string) (endRideSuccessful bool, err error) {
	var vehicle string
	row := database.QueryRow(BookedVehicleOfCustomerIdWithCertainStatus, customerId, config.StatusDriving)
	selectError := row.Scan(&vehicle)

	if selectError != nil {
		err = endingRideFailedError
	} else {
		_, updateError := database.Exec(UpdateBookingOnCustomerIdWithCertainStatus, config.StatusDone, customerId, config.StatusDriving)
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

func addInvoiceToBooking(bookingId uint32, invoiceId string) (successful bool, err error) {
	_, updateError := database.Exec(UpdateInvoiceOfBooking, invoiceId, bookingId, config.StatusDone)
	if updateError != nil {
		log.Log(updateError)
		err = commons.UnknownError
	} else {
		successful = true
	}

	return
}

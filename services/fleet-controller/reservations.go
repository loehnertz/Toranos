package main

import (
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
)

const AllReservations = "SELECT id, created_at, vehicle, customer, status FROM bookings WHERE status = $1 OR status = $2"
const AllUnbilledBookings = "SELECT id, created_at, customer, distance_driven, time_driven FROM bookings WHERE status = $1 AND invoice IS NULL"

func retrieveReservations() (reservations []*fleet_controller.RetrieveReservationsResponse_Reservation, err error) {
	rows, reservationRetrievalError := database.Query(AllReservations, config.StatusReserved, config.StatusDriving)
	defer rows.Close()
	if reservationRetrievalError != nil {
		log.Log(reservationRetrievalError)
		err = commons.UnknownError
	}

	for rows.Next() {
		var id uint32
		var createdAt time.Time
		var vehicle string
		var customer string
		var status int
		if rowsScanningError := rows.Scan(&id, &createdAt, &vehicle, &customer, &status); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = commons.UnknownError
		} else {
			reservations = append(reservations, &fleet_controller.RetrieveReservationsResponse_Reservation{
				Id:        id,
				CreatedAt: createdAt.Unix(),
				Vehicle:   vehicle,
				Customer:  customer,
				Status:    uint32(status),
			})
		}
	}

	if rowsError := rows.Err(); rowsError != nil {
		log.Log(rowsError)
		err = commons.UnknownError
	}

	return
}

func retrieveUnbilledBookings() (bookings []*fleet_controller.RetrieveUnbilledBookingsResponse_Booking, err error) {
	rows, bookingsRetrievalError := database.Query(AllUnbilledBookings, config.StatusDone)
	defer rows.Close()
	if bookingsRetrievalError != nil {
		log.Log(bookingsRetrievalError)
		err = commons.UnknownError
	}

	for rows.Next() {
		var id uint32
		var createdAt time.Time
		var customer string
		var distanceDriven uint32
		var timeDriven uint32
		if rowsScanningError := rows.Scan(&id, &createdAt, &customer, &distanceDriven, &timeDriven); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = commons.UnknownError
		} else {
			bookings = append(bookings, &fleet_controller.RetrieveUnbilledBookingsResponse_Booking{
				Id:             id,
				CreatedAt:      createdAt.Unix(),
				Customer:       customer,
				DistanceDriven: distanceDriven,
				TimeDriven:     timeDriven,
			})
		}
	}

	if rowsError := rows.Err(); rowsError != nil {
		log.Log(rowsError)
		err = commons.UnknownError
	}

	return
}

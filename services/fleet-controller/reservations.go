package main

import (
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
)

const AllReservations = "SELECT id, created_at, vehicle, customer, status FROM bookings WHERE status = $1 OR status = $2"

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

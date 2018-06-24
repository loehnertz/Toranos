package main

import (
	"database/sql"
	"github.com/loehnertz/toranos/src/commons"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
)

const AllReservations = "SELECT id, created_at, vehicle, customer FROM bookings WHERE status = $1"

func retrieveReservations(database *sql.DB) (reservations []*fleet_controller.RetrieveReservationsResponse_Reservation, err error) {
	rows, reservationRetrievalError := database.Query(AllReservations, config.ReservedStatus)
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
		if rowsScanningError := rows.Scan(&id, &createdAt, &vehicle, &customer); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = commons.UnknownError
		} else {
			reservations = append(reservations, &fleet_controller.RetrieveReservationsResponse_Reservation{
				Id:        id,
				CreatedAt: createdAt.Unix(),
				Vehicle:   vehicle,
				Customer:  customer,
			})
		}
	}

	if rowsError := rows.Err(); rowsError != nil {
		log.Log(rowsError)
		err = commons.UnknownError
	}

	return
}

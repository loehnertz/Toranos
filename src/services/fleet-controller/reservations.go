package main

import (
	"database/sql"
	"github.com/loehnertz/toranos/src/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
	"github.com/loehnertz/toranos/src/commons"
)

const ReservedStatus = 1

const AllReservations = "SELECT id, created_at, vehicle, customer, status FROM bookings WHERE status = $1"

func retrieveReservations(database *sql.DB) (reservations []*fleet_controller.RetrieveReservationsResponse_Reservation, err error) {
	rows, reservationRetrievalError := database.Query(AllReservations, ReservedStatus)
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
		var status uint32
		if rowsScanningError := rows.Scan(&id, &createdAt, &vehicle, &customer, &status); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = commons.UnknownError
		} else {
			reservations = append(reservations, &fleet_controller.RetrieveReservationsResponse_Reservation{
				Id:        id,
				CreatedAt: createdAt.Unix(),
				Vehicle:   vehicle,
				Customer:  customer,
				Status:    status,
			})
		}
	}
	if rowsError := rows.Err(); rowsError != nil {
		log.Log(rowsError)
		err = commons.UnknownError
	}

	return
}

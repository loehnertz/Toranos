package main

import (
	"github.com/loehnertz/toranos/common"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
)

const AllUnbilledBookings = "SELECT id, created_at, customer, distance_driven, time_driven FROM bookings WHERE (status = $1 OR status = $2) AND invoice IS NULL"
const AllBilledBookingsOfCustomer = "SELECT id, created_at, vehicle, customer, distance_driven, time_driven, invoice FROM bookings WHERE customer = $1 AND invoice IS NOT NULL"

func retrieveUnbilledBookings() (bookings []*fleet_controller.Booking, err error) {
	rows, bookingsRetrievalError := database.Query(
		AllUnbilledBookings,
		getStatusKeyByName(Done),
		getStatusKeyByName(Canceled),
	)
	defer rows.Close()
	if bookingsRetrievalError != nil {
		log.Log(bookingsRetrievalError)
		err = common.UnknownError
	}

	for rows.Next() {
		var id uint32
		var createdAt time.Time
		var customer string
		var distanceDriven uint32
		var timeDriven uint32
		if rowsScanningError := rows.Scan(
			&id,
			&createdAt,
			&customer,
			&distanceDriven,
			&timeDriven,
		); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = common.UnknownError
		} else {
			bookings = append(bookings, &fleet_controller.Booking{
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
		err = common.UnknownError
	}

	return
}

func retrieveBilledBookingsOfCustomer(customer string) (bookings []*fleet_controller.Booking, err error) {
	rows, bookingsRetrievalError := database.Query(AllBilledBookingsOfCustomer, customer)
	defer rows.Close()
	if bookingsRetrievalError != nil {
		log.Log(bookingsRetrievalError)
		err = common.UnknownError
	}

	for rows.Next() {
		var id uint32
		var createdAt time.Time
		var vehicle string
		var customer string
		var distanceDriven uint32
		var timeDriven uint32
		var invoice string
		if rowsScanningError := rows.Scan(
			&id,
			&createdAt,
			&vehicle,
			&customer,
			&distanceDriven,
			&timeDriven,
			&invoice,
		); rowsScanningError != nil {
			log.Log(rowsScanningError)
			err = common.UnknownError
		} else {
			bookings = append(bookings, &fleet_controller.Booking{
				Id:             id,
				CreatedAt:      createdAt.Unix(),
				VehicleId:      vehicle,
				Customer:       customer,
				DistanceDriven: distanceDriven,
				TimeDriven:     timeDriven,
				Invoice:        invoice,
			})
		}
	}

	if rowsError := rows.Err(); rowsError != nil {
		log.Log(rowsError)
		err = common.UnknownError
	}

	return
}

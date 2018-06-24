package config

// Service names
const FleetControllerName = "fleet-controller"
const FleetMonitorName = "fleet-monitor"

// Booking settings
const StatusError = 0
const StatusReserved = 1
const StatusDriving = 2
const StatusDone = 3
const ReservationTimeInSeconds = 900

// FleetMonitor intervals
const CheckForExpiredReservationsIntervalInSeconds = 60

package config

// Service names
const FleetControllerName = "fleet-controller"
const FleetMonitorName = "fleet-monitor"
const TelemetryName = "telemetry"

// Booking settings
const StatusError = 0
const StatusReserved = 1
const StatusDriving = 2
const StatusDone = 3
const ReservationTimeInSeconds = 900

// FleetMonitor cron intervals
const CheckForExpiredReservationsInterval = "0 */1 * * * *"

// Redis cache expiration times
const RedisAvailableVehiclesExpirationTimeInSeconds = 60

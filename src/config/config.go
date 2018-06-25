package config

// Service names
const FleetControllerName = "fleet-controller"
const FleetMonitorName = "fleet-monitor"
const TelemetryName = "telemetry"
const UserManagementName = "user-management"

// FleetController settings
const StatusError = 0
const StatusReserved = 1
const StatusDriving = 2
const StatusDone = 3
const ReservationTimeInSeconds = 900

// UserManagement settings
const TokenSecretTableKey = "token_secret"
const AudienceKeyCustomer = "customer"

// FleetMonitor cron intervals
const CheckForExpiredReservationsInterval = "0 */1 * * * *"

// Redis cache expiration times
const RedisAvailableVehiclesExpirationTimeInSeconds = 60

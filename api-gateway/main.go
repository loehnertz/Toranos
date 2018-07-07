package main

import (
	"github.com/gorilla/mux"
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-controller/proto"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/statistics/proto"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net/http"
	"time"
)

var service micro.Service
var fleetControllerService fleet_controller.FleetControllerService
var fleetMonitorService fleet_monitor.FleetMonitorService
var statisticsService statistics.StatisticsService
var userManagementService user_management.UserManagementService

func main() {
	// Create the service client
	service = micro.NewService(
		micro.Name(config.ApiGatewayName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	service.Init()
	serviceClient := service.Client()
	serviceClient.Init(client.Retries(3))

	// Register the router
	router := mux.NewRouter()

	// Create all the service clients
	initServiceClients(&serviceClient)

	// Create all the routes
	initRoutes(router)

	// Bind the server with the router to a port
	log.Fatal(http.ListenAndServe(config.ApiGatewayPort, router))
}

func initServiceClients(serviceClient *client.Client) {
	fleetControllerService = fleet_controller.NewFleetControllerService(config.FleetControllerName, *serviceClient)
	fleetMonitorService = fleet_monitor.NewFleetMonitorService(config.FleetMonitorName, *serviceClient)
	statisticsService = statistics.NewStatisticsService(config.StatisticsName, *serviceClient)
	userManagementService = user_management.NewUserManagementService(config.UserManagementName, *serviceClient)
}

func initRoutes(router *mux.Router) {
	// Public routes
	router.Handle("/register", http.HandlerFunc(registerNewUser)).Methods("POST")
	router.Handle("/login", http.HandlerFunc(getAuthToken)).Methods("POST")

	// Authenticated routes
	router.Handle("/available-vehicles", authenticationMiddleware(http.HandlerFunc(availableVehicles))).Methods("GET")
	router.Handle("/booking", authenticationMiddleware(http.HandlerFunc(createBooking))).Methods("POST")
	router.Handle("/booking", authenticationMiddleware(http.HandlerFunc(deleteBooking))).Methods("DELETE")
	router.Handle("/begin-ride", authenticationMiddleware(http.HandlerFunc(beginRide))).Methods("POST")
	router.Handle("/end-ride", authenticationMiddleware(http.HandlerFunc(endRide))).Methods("POST")
	router.Handle("/statistics", authenticationMiddleware(http.HandlerFunc(retrieveStatistics))).Methods("GET")
}

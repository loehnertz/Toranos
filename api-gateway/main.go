package main

import (
	"github.com/gorilla/mux"
	"github.com/loehnertz/toranos/config"
	"github.com/loehnertz/toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net/http"
	"time"
)

var service micro.Service
var fleetMonitor fleet_monitor.FleetMonitorService
var userManagement user_management.UserManagementService

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
	router.Use(authenticationMiddleware)

	// Create all the service clients
	initServiceClients(&serviceClient)

	// Create all the routes
	initRoutes(router)

	// Bind the server with the router to a port
	log.Fatal(http.ListenAndServe(config.ApiGatewayPort, router))
}

func initServiceClients(serviceClient *client.Client) {
	fleetMonitor = fleet_monitor.NewFleetMonitorService(config.FleetMonitorName, *serviceClient)
	userManagement = user_management.NewUserManagementService(config.UserManagementName, *serviceClient)
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/users/token", getAuthToken).Methods("GET")
	router.HandleFunc("/users/register", registerNewUser).Methods("POST")
	router.HandleFunc("/available-vehicles", availableVehicles).Methods("GET")
	router.HandleFunc("/booking", createBooking).Methods("POST")
	router.HandleFunc("/booking", deleteBooking).Methods("DELETE")
	router.HandleFunc("/begin-ride", beginRide).Methods("POST")
	router.HandleFunc("/end-ride", endRide).Methods("POST")
	router.HandleFunc("/statistics", retrieveStatistics).Methods("GET")
}

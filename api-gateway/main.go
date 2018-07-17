package main

import (
	"github.com/gorilla/mux"
	"github.com/loehnertz/Toranos/common"
	"github.com/loehnertz/Toranos/services/fleet-controller/proto"
	"github.com/loehnertz/Toranos/services/fleet-monitor/proto"
	"github.com/loehnertz/Toranos/services/statistics/proto"
	"github.com/loehnertz/Toranos/services/user-management/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net/http"
	"time"
)

var conf config.Config
var service micro.Service
var fleetControllerService fleet_controller.FleetControllerService
var fleetMonitorService fleet_monitor.FleetMonitorService
var statisticsService statistics.StatisticsService
var userManagementService user_management.UserManagementService

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Create the service client
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "api-gateway")),
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
	log.Fatal(http.ListenAndServe(common.GetConfigStringByPath(conf, "ports", "api-gateway"), router))
}

func initServiceClients(serviceClient *client.Client) {
	fleetControllerService = fleet_controller.NewFleetControllerService(common.GetConfigStringByPath(conf, "service-names", "fleet-controller"), *serviceClient)
	fleetMonitorService = fleet_monitor.NewFleetMonitorService(common.GetConfigStringByPath(conf, "service-names", "fleet-monitor"), *serviceClient)
	statisticsService = statistics.NewStatisticsService(common.GetConfigStringByPath(conf, "service-names", "statistics"), *serviceClient)
	userManagementService = user_management.NewUserManagementService(common.GetConfigStringByPath(conf, "service-names", "user-management"), *serviceClient)
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

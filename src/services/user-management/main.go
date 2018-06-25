package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/user-management/proto"
	"github.com/micro/go-micro"
	"time"
)

const DatabaseDriver = "postgres"
const DataSource = "user=jloehnertz dbname=toranos_users sslmode=disable"

var database *sql.DB
var service micro.Service
var tokenSigningKey []byte

type UserManagement struct{}

func (um *UserManagement) RegisterCustomer(ctx context.Context, req *user_management.RegisterCustomerRequest, res *user_management.RegisterCustomerResponse) error {
	registerSuccessful, token := registerCustomer(database, req)

	fmt.Println(registerSuccessful, token)

	return nil
}

func (um *UserManagement) IssueUserToken(ctx context.Context, req *user_management.IssueUserTokenRequest, res *user_management.IssueUserTokenResponse) error {
	issuingUserTokenSuccessful, token := issueUserToken(req.Email, req.Password)

	if issuingUserTokenSuccessful {
		res.Successful = true
		res.Token = token
	} else {
		res.Successful = false
	}

	return nil
}

func (um *UserManagement) AuthenticateUser(ctx context.Context, req *user_management.AuthenticateUserRequest, res *user_management.AuthenticateUserResponse) error {
	authenticated, email, role := authenticateUser(req.Token)

	if authenticated {
		res.Authenticated = true
		res.Email = email
		res.Role = role
	} else {
		res.Authenticated = false
	}

	return nil
}

func main() {
	// Connect the database
	var databaseError error
	database, databaseError = sql.Open(DatabaseDriver, DataSource)
	if databaseError != nil {
		panic(databaseError)
	}

	var retrieveTokenSecretError error
	tokenSigningKey, retrieveTokenSecretError = retrieveTokenSecret()
	if retrieveTokenSecretError != nil {
		panic(retrieveTokenSecretError)
	}

	// Create the service
	service = micro.NewService(
		micro.Name(config.UserManagementName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	user_management.RegisterUserManagementHandler(service.Server(), new(UserManagement))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

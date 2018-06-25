package main

import (
	"database/sql"
	"fmt"
	"github.com/loehnertz/toranos/src/config"
	"github.com/loehnertz/toranos/src/services/user-management/proto"
	"github.com/micro/go-log"
	"golang.org/x/crypto/bcrypt"
)

const InsertNewUser = "INSERT INTO users (email, password, token, first_name, last_name, license_id, license_verified) VALUES ($1, $2, $3, $4, $5, $6, $7)"

func registerCustomer(database *sql.DB, req *user_management.RegisterCustomerRequest) (successful bool, token string) {
	licenseVerified := verifyLicense(req.LicenseId)

	successfullyHashed, hash := hashAndSalt(req.Password)
	if !successfullyHashed {
		return
	}

	successfullyIssuedToken, tokenString := issueToken(req.Email, config.AudienceKeyCustomer)
	if !successfullyIssuedToken {
		return
	}

	_, insertError := database.Exec(InsertNewUser, req.Email, hash, tokenString, req.FirstName, req.LastName, req.LicenseId, licenseVerified)
	if insertError != nil {
		log.Log(insertError)
	} else {
		successful = true
		token = tokenString
	}

	return
}

func hashAndSalt(pwd string) (bool, string) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Log(err)
		return false, ""
	} else {
		return true, string(hashBytes)
	}
}

func verifyLicense(licenseId string) bool {
	fmt.Printf("The driver's license '%v' will be verified! \n", licenseId)
	return true // TODO: Connect to the "Driver's license" service
}

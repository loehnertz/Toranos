package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/loehnertz/toranos/src/config"
	"github.com/micro/go-log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const RetrieveTokenSecret = "SELECT value FROM secrets WHERE key = $1"

func issueToken(subject string, audience string) (bool, string) {
	// Create a new token
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"aud": audience,
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(tokenSigningKey))
	if err != nil {
		log.Log(err)
		return false, ""
	}
	return true, tokenString
}

func verifyToken(tokenString string) (successful bool, claims jwt.Claims) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the correct algorithm was used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSigningKey, nil
	})

	// Validate the token itself
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Log(err)
		successful = false
	} else if ok && token.Valid {
		successful = true
	}

	return
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		log.Log(err)
		return false
	} else {
		return true
	}
}

func retrieveTokenSecret() (string, error) {
	var tokenSecret string
	row := database.QueryRow(RetrieveTokenSecret, config.TokenSecretTableKey)
	selectError := row.Scan(&tokenSecret)
	if selectError != nil {
		log.Log(selectError)
		return "", selectError
	} else {
		return tokenSecret, nil
	}
}

package main

import (
	"context"
	"github.com/loehnertz/toranos/commons"
	"github.com/loehnertz/toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"net/http"
	"strings"
)

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")

		// Check if the user is authenticated
		resAuthenticateUser, errAuthenticateUser := userManagement.AuthenticateUser(context.TODO(), &user_management.AuthenticateUserRequest{Token: token[1]})

		if errAuthenticateUser != nil {
			log.Log(errAuthenticateUser)
			w.Write([]byte(commons.UnknownError.Error()))
		} else if resAuthenticateUser.Authenticated {
			// Call the next handler
			next.ServeHTTP(w, r)
		} else {
			w.Write([]byte(commons.NotAuthorizedError.Error()))
		}
	})
}

package main

import (
	"context"
	gorillacontext "github.com/gorilla/context"
	"github.com/loehnertz/Toranos/common"
	"github.com/loehnertz/Toranos/services/user-management/proto"
	"github.com/micro/go-log"
	"net/http"
	"strings"
)

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")

		// Check if the user is authenticated
		resAuthenticateUser, errAuthenticateUser := userManagementService.AuthenticateUser(context.TODO(), &user_management.AuthenticateUserRequest{Token: token[1]})

		if errAuthenticateUser != nil {
			log.Log(errAuthenticateUser)
			w.Write([]byte(common.UnknownError.Error()))
		} else if resAuthenticateUser.Authenticated {
			// Call the next handler
			gorillacontext.Set(r, "user", resAuthenticateUser)
			next.ServeHTTP(w, r)
		} else {
			w.Write([]byte(common.NotAuthorizedError.Error()))
		}
	})
}

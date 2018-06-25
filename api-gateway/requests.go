package main

type getAuthTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

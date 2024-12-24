package models

import "time"

type UserSessions struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResponseSignUp struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type ResponseSignIn struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

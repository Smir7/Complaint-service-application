package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Users struct {
	ID        uint      `db:"id" json:"id"`
	User_UUID uuid.UUID `db:"user_UUID" json:"user_UUID"`
	UserName  string    `db:"user_name" json:"user_name"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
	Role      Role      `db:"role" json:"role"`
}

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

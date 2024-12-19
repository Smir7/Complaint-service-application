package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id        int       `json:"id"`
	User_UUID uuid.UUID `json:"user_UUID"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
}

type UserSessions struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

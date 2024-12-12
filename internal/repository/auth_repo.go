package repository

import (
	"complaint_service/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	err := r.db.Get(&id, "SELECT COUNT(*) FROM users WHERE username = $1", user.Username)

	if id != 0 {

		return 0, fmt.Errorf("User %v already exists", user.Username)

	}

	query := fmt.Sprintf("INSERT INTO users (user_uuid,username,password,role) values($1,$2,$3,$4) RETURNING id")
	row := r.db.QueryRow(query, user.User_UUID, user.Username, user.Password, user.Role)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

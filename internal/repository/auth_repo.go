package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.UserSignUp) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres является конструктором структуры AuthPostgres. Принимает на вход переменную типа sqlx.DB и возвращает AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

/*
CreateUser отправляет INSERT запрос в базу данных для создания пользователя. Принимает на вход структуру User,
возвращает переменные id типа int и err типа error
*/
func (r *AuthPostgres) CreateUser(userModel models.UserSignUp) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	user := entity.User{
		Username:  userModel.Username,
		Password:  userModel.Password,
		User_UUID: userModel.User_UUID,
		Role:      userModel.Role,
	}
	query := fmt.Sprintf("INSERT INTO users (user_uuid,username,password,role) values($1,$2,$3,$4) RETURNING id")
	row := tx.QueryRow(query, user.User_UUID, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return id, nil
}

/*
GetUser отправляет SELECT запрос в базу данных для получения данных пользователя. Принимает на вход username и password,
возвращает структуру User и ошибку типа error
*/
func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

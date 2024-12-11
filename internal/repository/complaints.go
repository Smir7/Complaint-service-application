package repository

import "github.com/jmoiron/sqlx"

type ComplaintsRepository struct {
	Authorization
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{
		Authorization: NewAuthPostgres(db),
	}
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат

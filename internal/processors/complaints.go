package processors

import (
	"complaint_service/internal/repository"
)

type ComplaintsRepository interface {
	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	Authorization
}

// CreateComplaintsProcessor является конструктором структуры ComplaintsProcessor. Принимает на вход переменную типа sqlx.DB и возвращает ComplaintsProcessor
func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization: NewAuthService(complaintsRepository.Authorization),
	}
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository

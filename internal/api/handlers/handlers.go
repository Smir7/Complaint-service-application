package handlers

import (
	"complaint_service/internal/processors"

	"github.com/gofiber/fiber"
)

type ComplaintsProcessor interface {
	//имплиментируются методы из processors
}

type ComplaintsHandler struct {
	complaintsProcessor *processors.ComplaintsProcessor
}

// Функция CreateComplaintsHandler является конструктором структуры ComplaintsHandler. Принимает на вход переменную типа processors.ComplaintsProcessor и возвращает ComplaintsHandler.
func CreateComplaintsHandler(complaintsProcessor *processors.ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor: complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики

// Функция InitRoutes инициализирует роуты. Принимает на вход переменную типа fiber.App
func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
	app.Post("user/login", h.signIn)
}

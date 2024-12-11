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

func CreateComplaintsHandler(complaintsProcessor *processors.ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor: complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики

func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
}

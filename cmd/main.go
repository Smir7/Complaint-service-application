package main

import (
	"complaint_service/internal/api/handlers"
	"complaint_service/internal/processors"
	"complaint_service/internal/repository"
	"log/slog"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/log"
)

const port = ":8080"
const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	/*
		Далее передаем в наши ручки log *slog.Logger и с ним работаем.
		Для удобства, в каждой ручке можно использовать такую конструкцию, чтоб дальше подтягивалась информация.
		log := log.With(
			slog.String("где вылезла ошибка", op),
		)
		Инициализируем БД. И коннект прокидываем в CreateComplaintsRepository
		complaintsRepository := repository.CreateComplaintsRepository(db)

		Инициализируем ComplaintsProcessor где у нас будет бизнес логика
		complaintsProcessor := processors.CreateComplaintsProcessor(complaintsRepository)

		Инициализируем ComplaintsHandler, где у нас будут описаны хендлеры
		complaintsHandler := handlers.CreateComplaintsHandler(complaintsProcessor)
	*/

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Error("Create database error: %v", err)
	}

	repo := repository.CreateComplaintsRepository(db)
	service := processors.CreateComplaintsProcessor(repo)
	h := handlers.CreateComplaintsHandler(service)

	app := fiber.New()

	h.InitRoutes(app)

	/*
		Подключаем роуты. Прокидываем инициализированные хендлеры complaintsHandler
		routes.Complaints(app, complaintsHandler)
	*/
	log.Info("The server is running", slog.String("port", port))
	if err := app.Listen(port); err != nil {
		log.Error("Server startup error: %v", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log

}

package handlers

import (
	"complaint_service/internal/entity"

	"github.com/gofiber/fiber"
)

func (h *ComplaintsHandler) signUp(c *fiber.Ctx) {
	var input entity.User

	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.StatusBadRequest).SendString("wrong JSON")
		return
	}

	id, err := h.complaintsProcessor.Authorization.CreateUser(input)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}

	c.Status(fiber.StatusOK)
	c.JSONP(map[string]interface{}{
		"status":  "успешная регистрация",
		"user_id": id,
	})
}

package handlers

import (
	"complaint_service/internal/entity"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

const (
	successfulReg = "успешная регистрация"
	badRequest    = "incorrect request"
	serverError   = "ошибка севера"
)

func (h *ComplaintsHandler) signUp(c *fiber.Ctx) {
	var input entity.Users

	if err := c.BodyParser(&input); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSONP(
			entity.ResponseSignUp{
				Id:     0,
				Status: badRequest,
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	id, err := h.complaintsProcessor.Authorization.CreateUser(input)

	if err != nil {
		err = c.Status(fiber.StatusInternalServerError).JSONP(
			entity.ResponseSignUp{
				Id:     0,
				Status: fmt.Sprintf("%v: %v", serverError, err),
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		entity.ResponseSignUp{
			Id:     id,
			Status: successfulReg,
		})
	if err != nil {
		log.Println(err)
	}
}

func (h *ComplaintsHandler) signIn(c *fiber.Ctx) {
	var input entity.User

	if err := c.BodyParser(&input); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSONP(
			entity.ResponseSignIn{
				Token:  "",
				Status: badRequest,
			})
		if err != nil {
			log.Println(err)
		}
		return
	}
	token, err := h.complaintsProcessor.Authorization.GetToken(input.Username, input.Password)
	if err != nil {
		err = c.Status(fiber.StatusInternalServerError).JSONP(
			entity.ResponseSignIn{
				Token:  "",
				Status: fmt.Sprintf("%v: %v", serverError, err),
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		entity.ResponseSignIn{
			Token:  token,
			Status: successfulReg,
		})
	if err != nil {
		log.Println(err)
	}
}

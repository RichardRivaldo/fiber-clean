package controllers

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/dtos"
	"fiber-clean/src/repositories"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	query := dtos.Login{}
	if err := c.BodyParser(&query); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed logging in!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&query); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed logging in!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := repositories.Login(query.Email, query.Password)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusUnauthorized,
			Message: "Failed logging in!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully logged in!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func AddAuthRouter(router fiber.Router) {
	root := "/auth"
	router.Post(root, LoginHandler)
}

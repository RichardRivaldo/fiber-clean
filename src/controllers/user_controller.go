package controllers

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/dtos"
	"fiber-clean/src/models"
	"fiber-clean/src/repositories"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserHandler(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering user!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&user); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering user!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := repositories.RegisterUser(user)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed registering user!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully registered a user!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func AddUserRouter(router fiber.Router) {
	root := "/users"
	router.Post(root, RegisterUserHandler)
}

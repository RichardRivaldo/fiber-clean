package controllers

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/dtos"
	"fiber-clean/src/models"
	"fiber-clean/src/repositories"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminHandler(c *fiber.Ctx) error {
	admin := models.Admin{}

	if err := c.BodyParser(&admin); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering admin!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&admin); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering admin!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := repositories.RegisterAdmin(admin)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed registering admin!",
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

func GetStatisticsHandler(c *fiber.Ctx) error {
	result, err := repositories.GetStatistics()
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed retrieving statistics!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieving statistics!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	userId := c.Params("user_id")

	result, err := repositories.DeleteUser(userId)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed deleting user!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully deleted a user!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func AddAdminRouter(router fiber.Router) {
	root := "/admins"
	router.Get(root+"/statistics", GetStatisticsHandler)
	router.Post(root, RegisterAdminHandler)
	router.Delete(root+"/delete_user/:user_id", DeleteUserHandler)
}

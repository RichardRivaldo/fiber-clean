package controllers

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/dtos"
	"fiber-clean/src/middlewares"
	"fiber-clean/src/models"
	"fiber-clean/src/repositories"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func CreateNewCourseHandler(c *fiber.Ctx) error {
	course := models.Course{
		Exists: true,
	}

	if err := c.BodyParser(&course); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&course); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed registering course!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	image_file, _ := c.FormFile("image")
	if err := c.SaveFile(image_file, image_file.Filename); err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed registering course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	image_url, err := repositories.UploadImage(image_file.Filename)
	os.Remove(image_file.Filename)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed uploading image!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	course.Image = image_url
	result, err := repositories.CreateNewCourse(course)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed registering course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusCreated,
		Message: "Successfully registered a course!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusCreated).JSON(response)
}

func GetAllCoursesHandler(c *fiber.Ctx) error {
	result, err := repositories.GetAllCourses()
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed retrieving all courses!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieving all courses!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func UpdateCourseHandler(c *fiber.Ctx) error {
	course := models.Course{}
	if err := c.BodyParser(&course); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed updating course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&course); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed updating course!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	courseId := c.Params("course_id")
	result, err := repositories.UpdateCourse(courseId, course)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed updating course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully updated a course!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func DeleteCourseHandler(c *fiber.Ctx) error {
	courseId := c.Params("course_id")

	result, err := repositories.DeleteCourse(courseId)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed deleting course!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully deleted a course!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func SearchCourseHandler(c *fiber.Ctx) error {
	query := models.Query{}
	if err := c.BodyParser(&query); err != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed searching courses!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if validationErr := configs.Validator.Struct(&query); validationErr != nil {
		response := dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed searching courses!",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := repositories.QueryCourses(query.Filter, query.Projection, query.Sort)
	if err != nil {
		response := dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed searching courses!",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	response := dtos.Response{
		Status:  http.StatusOK,
		Message: "Successfully searching courses!",
		Data:    &fiber.Map{"data": result},
	}
	return c.Status(http.StatusOK).JSON(response)
}

func AddCourseRouter(router fiber.Router) {
	root := "/courses"
	router.Post(root, middlewares.IsAdmin(CreateNewCourseHandler))
	router.Put(root+"/:course_id", middlewares.IsAdmin(UpdateCourseHandler))
	router.Delete(root+"/:course_id", middlewares.IsAdmin(DeleteCourseHandler))

	router.Get(root, GetAllCoursesHandler)
	router.Post(root+"/search", SearchCourseHandler)
}

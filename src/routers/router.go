package routers

import (
	"fiber-clean/src/controllers"
	"fiber-clean/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App) {
	root := app.Group("api/v1")

	controllers.AddUserRouter(root)
	controllers.AddAuthRouter(root)

	app.Use(middlewares.AuthMiddleware())

	controllers.AddAdminRouter(root)
	controllers.AddCourseRouter(root)
}

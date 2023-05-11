package routers

import (
	"fiber-clean/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRouter(app *fiber.App) {
	root := app.Group("api/v1")

	controllers.AddUserRouter(root)
	controllers.AddAdminRouter(root)
	controllers.AddCourseRouter(root)
	controllers.AddAuthRouter(root)
}

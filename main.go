package main

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.CreateDBConnection()
	routers.SetRouter(app)

	app.Listen(":8000")
}

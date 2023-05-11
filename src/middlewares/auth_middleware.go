package middlewares

import (
	"fiber-clean/src/configs"
	"fiber-clean/src/dtos"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func jwtError(c *fiber.Ctx, err error) error {
	status := http.StatusUnauthorized
	if err.Error() == "Missing or malformed JWT" {
		status = http.StatusBadRequest
	}

	response := dtos.Response{
		Status:  status,
		Message: err.Error(),
		Data:    nil,
	}
	return c.Status(status).JSON(response)
}

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(configs.GetEnv("JWT_AUTH_SECRET")),
		ErrorHandler: jwtError,
	})
}

func IsAdmin(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["is_admin"].(bool)

		if !isAdmin {
			response := dtos.Response{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized!",
				Data:    nil,
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}
		return next(c)
	}
}

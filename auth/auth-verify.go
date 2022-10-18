package auth

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthVerify(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		sub, err := GetSubFromToken(tokenString)
		if err != nil {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		a := context.WithValue(context.Background(), UserSubKey, sub)
		c.SetUserContext(a)

		return c.Next()
	})
}

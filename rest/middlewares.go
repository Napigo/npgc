package rest

import (
	"context"
	"strings"

	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/jwt"
)

type userSubKey int

var (
	UserSubKey userSubKey
	JWTSecret  = []byte("NAPIGO-JWT-SECRET")
)

func UserSubMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		sub, err := getSubFromToken(tokenString)
		if err != nil {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		a := context.WithValue(context.Background(), UserSubKey, sub)
		c.SetUserContext(a)

		return c.Next()
	})
}

func getSubFromToken(tokenString string) (string, error) {
	npglogger.Info("getSubFromToken" + tokenString)
	byteToken := []byte(tokenString)

	verifiedToken, err := jwt.VerifyWithHeaderValidator(jwt.HS256, JWTSecret, byteToken, func(alg string, headerDecoded []byte) (jwt.Alg, jwt.PublicKey, jwt.InjectFunc, error) {
		return jwt.HS256, JWTSecret, nil, nil
	})
	if err != nil {
		return "", nil
	}

	sub := verifiedToken.StandardClaims.Subject
	return sub, nil
}

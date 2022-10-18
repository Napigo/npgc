package auth

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/jwt"
)

func AuthVerify(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		sub, err := _getSubFromToken(tokenString)
		if err != nil {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		a := context.WithValue(context.Background(), UserSubKey, sub)
		c.SetUserContext(a)

		return c.Next()
	})
}

// Util function to extact the claim "subject" from
// the jwt token string value
func _getSubFromToken(tokenString string) (string, error) {
	npglogger.Info("getSubFromToken" + tokenString)
	byteToken := []byte(tokenString)

	secret := os.Getenv("JWT_SECRETS")

	if len(secret) == 0 {
		return "", errors.New("jwt secrets not found")
	}

	verifiedToken, err := jwt.VerifyWithHeaderValidator(jwt.HS256, secret, byteToken, func(alg string, headerDecoded []byte) (jwt.Alg, jwt.PublicKey, jwt.InjectFunc, error) {
		return jwt.HS256, secret, nil, nil
	})
	if err != nil {
		return "", err
	}

	sub := verifiedToken.StandardClaims.Subject
	return sub, nil
}

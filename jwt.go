package npgc

import (
	"errors"
	"os"
	"time"

	"github.com/kataras/jwt"
)

// Use for our napig custom jwt headers throughout
// all microservices.
type JWTHeader struct {
	Typ string `json:"typ"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
}

type JWTBuilder struct {
	Expiry   int64
	IssuedAt int64
	Issuer   string
	Subject  string
	Secret   []byte
}

func (jt JWTBuilder) CreateJWT() (*string, error) {
	standardClaims := jwt.Claims{
		Expiry:   jt.Expiry,
		IssuedAt: jt.IssuedAt,
		Issuer:   jt.Issuer,
		Subject:  jt.Subject,
	}

	header := JWTHeader{
		Typ: "JWT",
		Kid: os.Getenv("JWT_KID"),
		Alg: jwt.HS256.Name(),
	}
	token, err := jwt.SignWithHeader(jwt.HS256, jt.Secret, standardClaims, header, jwt.MaxAge(60*time.Minute))
	if err != nil {
		return nil, err
	}

	stringToken := string(token)
	return &stringToken, nil
}

func GetSubFromToken(tokenString string) (string, error) {
	byteToken := []byte(tokenString)

	secret := []byte(os.Getenv("JWT_SECRETS"))

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

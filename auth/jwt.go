package auth

import (
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

type JWTToken struct {
	Expiry   int64
	IssuedAt int64
	Issuer   string
	Subject  string
	Kid      string
	Typ      string
	Alg      string
	Secret   []byte
}

func (jt JWTToken) CreateJWT() (*string, error) {
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

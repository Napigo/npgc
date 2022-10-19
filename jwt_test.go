package npgc_test

import (
	"os"
	"testing"
	"time"

	npgc "github.com/Napigo/npgc"
)

type TPayload struct {
	UserId string
	Result string
}

func PrepareTestSuites() {
	os.Setenv("JWT_SECRETS", "NAPIGO-JWT-SECRETS")
	os.Setenv("JWT_KID", "NAPIGO-JWT_KEY")
	os.Setenv("JWT_ISSUER", "napigo")
}

func Test_GetSubFromToken(t *testing.T) {
	os.Setenv("JWT_SECRETS", "NAPIGO-JWT-SECRETS")
	os.Setenv("JWT_KID", "NAPIGO-JWT-KEY")
	os.Setenv("JWT_ISSUER", "napigo")
	now := time.Now()

	testPayload := TPayload{
		UserId: "Standard-A",
		Result: "Standard-A",
	}

	jToken := npgc.JWTBuilder{
		Expiry:   now.Add(15 * time.Minute).Unix(),
		IssuedAt: now.Unix(),
		Issuer:   os.Getenv("JWT_ISSUER"),
		Subject:  testPayload.UserId,
		Secret:   []byte(os.Getenv("JWT_SECRETS")),
	}

	token, err := jToken.CreateJWT()
	if err != nil {
		t.Error("Failed to generate token from CreateJWT()")
	}
	sToken := *token

	sub, err := npgc.GetSubFromToken(sToken)
	if err != nil {
		t.Error("Failed to extract subject ...GetSubFromToken() ")
	}
	if sub != testPayload.Result {
		t.Errorf("Failed, subject is not expected, value [%s]", sub)
	}
}

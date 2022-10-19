package npgc_test

import (
	"log"
	"testing"

	npgc "github.com/Napigo/npgc"
)

func TestConfig(t *testing.T) {
	npgc.Load("./test/.env.test")
	config := npgc.Config
	log.Println(config)
}

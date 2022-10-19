package npgc

import (
	"log"

	"github.com/joho/godotenv"
)

// This function will load the env vars form the .env, so
// all repo uses this function must provide .env file in their
// directory
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file, please check if you have .env in the project dir")
	}
	log.Print("Environment loaded.")
}

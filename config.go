package npgc

import (
	"os"

	"github.com/Napigo/npglogger"
	"github.com/joho/godotenv"
)

type (
	AppConfig struct {
		Port        string
		ServiceName string
		DB          DBConfig
		Auth        AuthConfig
	}
	DBConfig struct {
		URI      string
		AuthType string
		Username string
		Password string
	}

	AuthConfig struct {
		JWTSecret      string
		PublicKeyPath  string
		PrivateKeyPath string
		Kid            string
		Issuer         string
	}
)

var Config *AppConfig

func Load(fileName string) {
	if err := godotenv.Load(fileName); err != nil {
		npglogger.Fatal("Failed to load .env files")
	}
	Config = New()
}

func New() *AppConfig {
	return &AppConfig{
		Port:        getStringEnv("SERVICE_PORT", ":80"),
		ServiceName: getStringEnv("SERVICE_NAME", "untitled-service"),
		DB: DBConfig{
			URI:      getStringEnv("MONGO_DATABASE_URI", "mongodb://localhost:27017"),
			AuthType: getStringEnv("MONGO_AUTH_TYPE", "none"),
			Username: getStringEnv("MONGO_APP_USERNAME", ""),
			Password: getStringEnv("MONGO_APP_PASSWORD", ""),
		},
		Auth: AuthConfig{
			JWTSecret:      getStringEnv("JWT_SECRETS", ""),
			PublicKeyPath:  getStringEnv("JWT_PUBLIC_KEY_PATH", ""),
			PrivateKeyPath: getStringEnv("JWT_PRIVATE_KEY_PATH", ""),
			Kid:            getStringEnv("JWT_KID", "napigo-kid"),
			Issuer:         getStringEnv("JWT_ISSUER", "napigo"),
		},
	}
}

func getStringEnv(key string, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}

// Will be using below functions soon when required

// func GetEnvAsInt(name string, defaultVal int) int {
// 	valueStr := getStringEnv(name, "")
// 	if value, err := strconv.Atoi(valueStr); err == nil {
// 		return value
// 	}

// 	return defaultVal
// }

// func getEnvAsBool(name string, defaultVal bool) bool {
// 	valStr := getStringEnv(name, "")
// 	if val, err := strconv.ParseBool(valStr); err == nil {
// 		return val
// 	}

// 	return defaultVal
// }

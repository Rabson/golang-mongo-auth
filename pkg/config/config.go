package config

import (
	"os"
)

func GetEnv(key, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		println("Using environment variable:", key, "=", value)
		return value
	}
	return defaultValue
}

func GetPort() string {
	return GetEnv("PORT", "8080")
}
func GetMongoURI() string {
	return GetEnv("MONGO_URI", "")
}
func GetDbName() string {
	return GetEnv("DB_NAME", "")
}
func GetJwtSecrets() string {
	return GetEnv("JWT_SECRET", "")
}

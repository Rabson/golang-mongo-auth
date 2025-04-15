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

func GetS3Region() string {
	return GetEnv("AWS_REGION", "")
}
func GetS3AccessKeyId() string {
	return GetEnv("AWS_ACCESS_KEY_ID", "")
}
func GetS3SecretAccessKey() string {
	return GetEnv("AWS_SECRET_ACCESS_KEY", "")
}

func GetS3RBucket() string {
	return GetEnv("BUCKET_NAME", "")
}
func GetS3Endpoint() string {
	return "https://" + GetEnv("AWS_REGION", "") + ".linodeobjects.com"
}

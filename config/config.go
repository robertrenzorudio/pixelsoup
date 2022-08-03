package config

import (
	"os"
	"strings"
)

type Config struct {
	Env           string
	JobsQueueName string
}

func New() *Config {
	return &Config{
		Env:           getEnvAsEnvType("ENV", "development"),
		JobsQueueName: getEnv("AWS_JOBS_QUEUE_NAME", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsEnvType(key string, defaultVal string) string {
	e := getEnv(key, defaultVal)
	e = strings.ToLower(e)

	if e == "dev" || e == "development" || e == "prod" || e == "production" {
		return e
	}

	return defaultVal
}

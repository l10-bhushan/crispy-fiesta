package config

import (
	"os"
)

// We are separating all the components of our applications because we will be implementing
// dependency injection

// We create a config struct that will hold our databaseurl and port for the server
type Config struct {
	DatabaseURL string
	HTTPPort    string
}

// Load function that will load env variables into config struct and return us Config
func Load() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		HTTPPort:    getEnv("HTTPPORT", "8080"),
	}
}

// We are using this as a utlity function, where we provide the key and a fallback
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		value = fallback
	}

	return value
}

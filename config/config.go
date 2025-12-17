package config

import "os"

type Config struct {
	DBUrl string
}

func Load() *Config {
	return &Config{
		DBUrl: getEnv("DATABASE_URL", "postgres://postgres:sanuj2004@localhost:5432/ainyx_db?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

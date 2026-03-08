package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	StoragePath string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://xstorage:xstorage@localhost:5432/xstorage?sslmode=disable"),
		StoragePath: getEnv("STORAGE_PATH", "/data"),
		JWTSecret:   getEnv("JWT_SECRET", "change-me-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

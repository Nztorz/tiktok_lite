package app

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbURL       string
	environment string
	PORT        string
}

func LoadConfig() Config {
	godotenv.Load()
	return Config{
		dbURL:       os.Getenv("DB_URL"),
		environment: os.Getenv("environment"),
		PORT:        os.Getenv("PORT"),
	}
}

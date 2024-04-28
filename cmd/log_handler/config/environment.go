package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}

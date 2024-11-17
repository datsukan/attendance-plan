package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := loadEnv(); err != nil {
		panic(err)
	}
}

func loadEnv() error {
	if os.Getenv("ENV") == "prd" {
		return nil
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}

	return nil
}

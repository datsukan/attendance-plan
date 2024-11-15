package infrastructure

import "github.com/joho/godotenv"

func init() {
	if err := loadEnv(); err != nil {
		panic(err)
	}
}

func loadEnv() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}

	return nil
}

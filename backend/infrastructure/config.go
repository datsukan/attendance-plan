package infrastructure

import (
	"os"
	"strconv"
)

var config Config

type Config struct {
	SecretKey     string
	TokenLifeTime int
}

func init() {
	secretKey := os.Getenv("SESSION_SECRET_KEY")
	tokenLifeTimeStr := os.Getenv("SESSION_TOKEN_LIFE_TIME")
	tokenLifeTime, err := strconv.Atoi(tokenLifeTimeStr)
	if err != nil {
		tokenLifeTime = 30
	}

	config = Config{
		SecretKey:     secretKey,
		TokenLifeTime: tokenLifeTime,
	}
}

func GetConfig() Config {
	return config
}

func GetSecretKey() string {
	return config.SecretKey
}

func GetTokenLifeTime() int {
	return config.TokenLifeTime
}

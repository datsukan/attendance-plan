package infrastructure

import (
	"os"
	"strconv"
	"strings"
)

var config Config

type Config struct {
	ServiceName   string
	BaseUrl       string
	SecretKey     string
	TokenLifeDays int
	SESRegion     string
	SenderEmail   string
	SenderName    string
	AdminEmails   []string
}

func init() {
	serviceName := os.Getenv("SERVICE_NAME")
	baseUrl := os.Getenv("BASE_URL")
	secretKey := os.Getenv("SESSION_SECRET_KEY")
	tokenLifeDaysStr := os.Getenv("SESSION_TOKEN_LIFE_DAYS")
	tokenLifeDays, err := strconv.Atoi(tokenLifeDaysStr)
	if err != nil {
		tokenLifeDays = 30
	}

	sesRegion := os.Getenv("SES_REGION")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderName := os.Getenv("SENDER_NAME")

	var adminEmails []string
	for _, e := range strings.Split(os.Getenv("ADMIN_EMAILS"), ",") {
		if e = strings.TrimSpace(e); e != "" {
			adminEmails = append(adminEmails, e)
		}
	}

	config = Config{
		ServiceName:   serviceName,
		BaseUrl:       baseUrl,
		SecretKey:     secretKey,
		TokenLifeDays: tokenLifeDays,
		SESRegion:     sesRegion,
		SenderEmail:   senderEmail,
		SenderName:    senderName,
		AdminEmails:   adminEmails,
	}
}

func GetConfig() Config {
	return config
}

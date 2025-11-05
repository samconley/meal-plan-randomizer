package model

import (
	"os"
	"strconv"
	"strings"
)

type SmsConfig struct {
	FromEmail     string
	ToList        []string
	EmailPassword string
	SmtpHost      string
	SmtpPort      string
}

func NewSmsConfig() *SmsConfig {
	return &SmsConfig{
		FromEmail:     os.Getenv("FROM_EMAIL"),
		ToList:        strings.Split(os.Getenv("TO_LIST"), ","),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
		SmtpHost:      os.Getenv("SMTP_HOST"),
		SmtpPort:      os.Getenv("SMTP_HOST_PORT"),
	}
}

type MealConfig struct {
	SourceFileDir       string
	NumberOfMealsToSend int
}

func NewMealConfig() (*MealConfig, error) {
	numMealsToSendStr := os.Getenv("NUM_MEALS_TO_SEND")
	numMealsToSend, err := strconv.Atoi(numMealsToSendStr)
	if err != nil {
		return nil, err
	}

	return &MealConfig{
		SourceFileDir:       os.Getenv("SOURCE_FILE_DIR"),
		NumberOfMealsToSend: numMealsToSend,
	}, nil
}

package service

import (
	"log"
	"meal-plan-randomizer/internal/model"
	"net/smtp"
	"strings"
)

type SmsMessageService struct {
	Config *model.SmsConfig
}

func NewSmsMessageService(cfg *model.SmsConfig) *SmsMessageService {
	return &SmsMessageService{
		Config: cfg,
	}
}

func (s *SmsMessageService) ComposeMessageFromMeal(meal *model.Meal) {
	sb := &strings.Builder{}

	sb.WriteString("\n" + meal.Name)

	sb.WriteString("\nSupplies - ")
	for _, ingredient := range meal.Ingredients {
		sb.WriteString(ingredient + ",")
	}

	if len(meal.Sides) > 0 {
		sb.WriteString("\nSides - ")

		for _, side := range meal.Sides {
			sb.WriteString(side + ",")
		}
	}

	s.sendMsg(sb.String())
}

func (s *SmsMessageService) sendMsg(msg string) {
	cfg := s.Config
	auth := smtp.PlainAuth("", cfg.FromEmail, cfg.EmailPassword, cfg.SmtpHost)

	if err := smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.FromEmail, cfg.ToList, []byte(msg)); err != nil {
		log.Printf("error: %v", err)
	}
}

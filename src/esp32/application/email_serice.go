package application

import (
    "API_ONE/src/esp32/domain/entities"
    "API_ONE/src/esp32/domain/repositories"
)

type EmailService struct {
    emailRepo repositories.EmailRepository
}

func NewEmailService(emailRepo repositories.EmailRepository) *EmailService {
    return &EmailService{emailRepo: emailRepo}
}

func (s *EmailService) EnviarEmail(to, subject, body string) error {
    email := entities.Email{
        To:      to,
        Subject: subject,
        Body:    body,
    }

    return s.emailRepo.EnviarEmail(email)
}
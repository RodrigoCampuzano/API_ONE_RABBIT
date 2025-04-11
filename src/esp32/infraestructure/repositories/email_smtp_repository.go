package repositories

import (
	"log"
	"API_ONE/src/esp32/domain/entities"
)

type EmailMockRepository struct{}

func NewEmailMockRepository() *EmailMockRepository {
	return &EmailMockRepository{}
}

func (r *EmailMockRepository) EnviarEmail(email entities.Email) error {
	log.Printf("\n📨 Email simulado enviado:\nPara: %s\nAsunto: %s\nMensaje: %s\n", 
		email.To, email.Subject, email.Body)
	return nil
}
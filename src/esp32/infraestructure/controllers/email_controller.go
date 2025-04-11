package controllers

import (
	"encoding/json"
	"net/http"

	"API_ONE/src/esp32/application"
)

type EmailController struct {
	emailService *application.EmailService
}

func NewEmailController(emailService *application.EmailService) *EmailController {
	return &EmailController{emailService: emailService}
}

func (c *EmailController) EnviarEmail(w http.ResponseWriter, r *http.Request) {
	var data struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	
	json.NewDecoder(r.Body).Decode(&data)
	
	c.emailService.EnviarEmail(data.To, data.Subject, data.Body)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Correo procesado",
	})
}
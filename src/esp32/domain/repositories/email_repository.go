package repositories

import "API_ONE/src/esp32/domain/entities"

type EmailRepository interface {
    EnviarEmail(email entities.Email) error
}
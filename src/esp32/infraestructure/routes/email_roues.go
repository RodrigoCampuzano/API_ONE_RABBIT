package routes

import (
    "net/http"

    "API_ONE/src/esp32/application"
    "API_ONE/src/esp32/domain/repositories"
    "API_ONE/src/esp32/infraestructure/controllers"
)

func NewEmailRouter(emailRepo repositories.EmailRepository) http.Handler {
    mux := http.NewServeMux()
    emailService := application.NewEmailService(emailRepo)
    emailController := controllers.NewEmailController(emailService)
    
    // IMPORTANTE: Usar rutas relativas vacías
    mux.HandleFunc("/", emailController.EnviarEmail)
    return mux
}
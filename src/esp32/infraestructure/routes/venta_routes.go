package routes

import (
    "net/http"

    "API_ONE/src/esp32/application"
    "API_ONE/src/esp32/domain/repositories"
    "API_ONE/src/esp32/infraestructure/controllers"
)

func NewVentaRouter(ventaRepo repositories.VentaRepository, emailRepo repositories.EmailRepository) http.Handler {
    mux := http.NewServeMux()
    emailService := application.NewEmailService(emailRepo)
    ventaService := application.NewVentaService(ventaRepo, emailService)
    ventaController := controllers.NewVentaController(ventaService)
    
    // IMPORTANTE: Usar rutas relativas vacías
    mux.HandleFunc("/", ventaController.EnviarVenta)
    return mux
}
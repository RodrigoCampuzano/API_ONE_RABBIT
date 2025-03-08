package routes

import (
	"net/http"

	"API_ONE/src/esp32/application"
	"API_ONE/src/esp32/infraestructure/controllers"
	"API_ONE/src/esp32/infraestructure/repositories"
)

func NewVentaRouter() http.Handler {
	mux := http.NewServeMux()

	ventaRepo := repositories.NewVentaRepositoryRabbitMQ()
	ventaService := application.NewVentaService(ventaRepo)
	ventaController := controllers.NewVentaController(ventaService)

	mux.HandleFunc("/enviar-venta", ventaController.EnviarVenta)

	return mux
}
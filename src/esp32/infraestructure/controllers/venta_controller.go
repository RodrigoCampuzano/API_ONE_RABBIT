package controllers

import (
	"encoding/json"
	"net/http"

	"API_ONE/src/esp32/application"
)

type VentaController struct {
	ventaService *application.VentaService
}

func NewVentaController(ventaService *application.VentaService) *VentaController {
	return &VentaController{ventaService: ventaService}
}

func (c *VentaController) EnviarVenta(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	producto := data["producto"].(string)
	cantidad := int(data["cantidad"].(float64))

	err := c.ventaService.ProcesarVenta(producto, cantidad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Venta recibida y enviada a RabbitMQ"))
}
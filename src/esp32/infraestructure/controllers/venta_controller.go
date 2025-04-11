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
	// Verificar que el cuerpo no esté vacío
	if r.Body == nil {
		http.Error(w, "El cuerpo de la petición no puede estar vacío", http.StatusBadRequest)
		return
	}

	// Decodificar el JSON
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error al decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validar y obtener el campo 'producto'
	productoInterface, exists := data["producto"]
	if !exists {
		http.Error(w, "El campo 'producto' es requerido", http.StatusBadRequest)
		return
	}

	producto, ok := productoInterface.(string)
	if !ok {
		http.Error(w, "El campo 'producto' debe ser una cadena de texto", http.StatusBadRequest)
		return
	}

	// Validar y obtener el campo 'cantidad'
	cantidadInterface, exists := data["cantidad"]
	if !exists {
		http.Error(w, "El campo 'cantidad' es requerido", http.StatusBadRequest)
		return
	}

	cantidadFloat, ok := cantidadInterface.(float64)
	if !ok {
		http.Error(w, "El campo 'cantidad' debe ser un número", http.StatusBadRequest)
		return
	}
	cantidad := int(cantidadFloat)

	// Procesar la venta
	err := c.ventaService.ProcesarVenta(producto, cantidad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Venta recibida y enviada a RabbitMQ",
	})
}
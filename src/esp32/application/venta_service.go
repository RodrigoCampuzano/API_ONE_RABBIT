package application

import (
	"API_ONE/src/esp32/domain/entities"
	"API_ONE/src/esp32/domain/repositories"
	"fmt"
)

type VentaService struct {
	ventaRepo repositories.VentaRepository
	emailService *EmailService
}

func NewVentaService(ventaRepo repositories.VentaRepository, emailService *EmailService) *VentaService {
	return &VentaService{
		ventaRepo: ventaRepo,
		emailService: emailService,
	}
}

func (s *VentaService) ProcesarVenta(producto string, cantidad int) error {
	venta := entities.Venta{
		Producto: producto,
		Cantidad: cantidad,
	}

	err := s.ventaRepo.EnviarVenta(venta)
	if err != nil {
		return err
	}

	emailTo := "notificaciones@empresa.com" 
	emailSubject := "Nueva venta procesada"
	emailBody := fmt.Sprintf("Se ha procesado una nueva venta:\nProducto: %s\nCantidad: %d", producto, cantidad)
	
	return s.emailService.EnviarEmail(emailTo, emailSubject, emailBody)
}
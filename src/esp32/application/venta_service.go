package application

import (
	"API_ONE/src/esp32/domain/entities"
	"API_ONE/src/esp32/domain/repositories"
)

type VentaService struct {
	ventaRepo repositories.VentaRepository
}

func NewVentaService(ventaRepo repositories.VentaRepository) *VentaService {
	return &VentaService{ventaRepo: ventaRepo}
}

func (s *VentaService) ProcesarVenta(producto string, cantidad int) error {
	venta := entities.Venta{
		Producto: producto,
		Cantidad: cantidad,
	}

	return s.ventaRepo.EnviarVenta(venta)
}

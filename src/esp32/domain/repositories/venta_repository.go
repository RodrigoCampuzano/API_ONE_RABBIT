package repositories

import "API_ONE/src/esp32/domain/entities"

type VentaRepository interface {
	EnviarVenta(venta entities.Venta) error
}
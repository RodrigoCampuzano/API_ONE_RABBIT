package repositories

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"API_ONE/src/esp32/domain/entities"
	"API_ONE/src/esp32/domain/repositories"
)

type VentaRepositoryRabbitMQ struct{}

func NewVentaRepositoryRabbitMQ() repositories.VentaRepository {
	return &VentaRepositoryRabbitMQ{}
}

func (r *VentaRepositoryRabbitMQ) EnviarVenta(venta entities.Venta) error {
	conn, err := amqp.Dial("amqp://admin:password@52.0.68.153:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"cola_ventas",
		false, false, false, false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	body, err := json.Marshal(venta)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	return nil
}

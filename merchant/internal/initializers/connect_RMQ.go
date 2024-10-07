package initializers

import (
	"context"
	"log"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

func ConnectRMQ(config *Config, ctx context.Context) *amqp.Connection {
	rabbitConfig := rabbitmq.RabbitMQConfig{
		Host:     config.AMQP_SERVER_HOST,
		Port:     config.AMQP_SERVER_PORT,
		User:     config.AMQP_SERVER_USER,
		Password: config.AMQP_SERVER_PASSWORD,
	}
	RMQ, err := rabbitmq.NewRabbitMQConn(&rabbitConfig, ctx)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	return RMQ
}

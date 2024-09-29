package rabbit_handler

import (
	"encoding/json"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models/dto"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type PaymentDependencies struct {
	OrderService *services.OrderService
	Logger       zerolog.Logger
}

func PaymentOrderCompleted(queue string, msg amqp.Delivery, dependencies *PaymentDependencies) error {
	dependencies.Logger.Info().Msgf("Message received on queue: %s with message: %s", queue, string(msg.Body))

	var paymentOrderCompleted dto.PaymentResponse

	err := json.Unmarshal(msg.Body, &paymentOrderCompleted)
	if err != nil {
		return err
	}

	_, err = dependencies.OrderService.PaymentOrderCompleted(paymentOrderCompleted)
	return err
}

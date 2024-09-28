package rabbit_handler

import (
	"encoding/json"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type CreateOrderSucceedDependencies struct {
	InventoryService *services.InventoryService
	Logger           zerolog.Logger
}

func CreateOrderSucceed(queue string, msg amqp.Delivery, dependencies *CreateOrderSucceedDependencies) error {
	dependencies.Logger.Info().Msgf("Message received on queue: %s with message: %s", queue, string(msg.Body))

	var orderSucceed dto.CreateOrderSucceed

	err := json.Unmarshal(msg.Body, &orderSucceed)
	if err != nil {
		return err
	}

	err = dependencies.InventoryService.CreateOrderSucceed(&orderSucceed)
	return err
}

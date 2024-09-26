package rabbitmq

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//go:generate mockery --name IPublisher
type IPublisher interface {
	PublishMessage(msg interface{}) error
}

type Publisher struct {
	cfg          *RabbitMQConfig
	conn         *amqp.Connection
	log          zerolog.Logger
	ctx          context.Context
	exchangeName string
	exchangeType string
	queueName    string
}

func (p Publisher) PublishMessage(msg interface{}) error {

	data, err := jsoniter.Marshal(msg)

	if err != nil {
		p.log.Fatal().Err(err).Msg("Error in marshalling message to publish message")
		return err
	}

	// Inject the context in the headers
	// headers := otel.InjectAMQPHeaders(ctx)

	channel, err := p.conn.Channel()
	if err != nil {
		p.log.Fatal().Err(err).Msg("Error in opening channel to consume message")
		return err
	}

	defer channel.Close()

	err = channel.ExchangeDeclare(
		p.exchangeName, // name
		p.exchangeType, // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		p.log.Fatal().Err(err).Msg("Error in declaring exchange to publish message")
		return err
	}

	// TODO
	correlationId := ""

	publishingMsg := amqp.Publishing{
		Body:          data,
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		MessageId:     uuid.NewV4().String(),
		Timestamp:     time.Now(),
		CorrelationId: correlationId,
	}

	err = channel.Publish(p.exchangeName, p.queueName, false, false, publishingMsg)

	if err != nil {
		p.log.Fatal().Err(err).Msg("Error in publishing message")
		return err
	}

	p.log.Info().Msgf("Published message: %s", publishingMsg.Body)

	return nil
}

func NewPublisher(
	ctx context.Context,
	cfg *RabbitMQConfig,
	conn *amqp.Connection,
	log zerolog.Logger,
	exchangeName string,
	exchangeType string,
	queueName string,
) IPublisher {
	return &Publisher{
		ctx:          ctx,
		cfg:          cfg,
		conn:         conn,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queueName:    queueName,
	}
}

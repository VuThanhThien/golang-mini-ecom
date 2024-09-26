package rabbitmq

import (
	"context"
	"reflect"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

//go:generate mockery --name IConsumer
type IConsumer[T any] interface {
	ConsumeMessage(msg interface{}, dependencies T) error
	IsConsumed(msg interface{}) bool
}

var consumedMessages []string

type Consumer[T any] struct {
	cfg          *RabbitMQConfig
	conn         *amqp.Connection
	log          zerolog.Logger
	handler      func(queue string, msg amqp.Delivery, dependencies T) error
	ctx          context.Context
	exchangeName string
	exchangeType string
	queueName    string
	routingKey   string
}

func (c Consumer[T]) ConsumeMessage(msg interface{}, dependencies T) error {
	ch, err := c.conn.Channel()
	if err != nil {
		c.log.Fatal().Err(err).Msg("Error in opening channel to consume message")
		return err
	}

	err = ch.ExchangeDeclare(
		c.exchangeName, // name
		c.exchangeType, // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		c.log.Fatal().Err(err).Msg("Error in declaring exchange to consume message")
		return err
	}

	q, err := ch.QueueDeclare(
		c.queueName, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		c.log.Fatal().Err(err).Msg("Error in declaring queue to consume message")
		return err
	}

	err = ch.QueueBind(
		q.Name,         // queue name
		c.routingKey,   // routing key
		c.exchangeName, // exchange
		false,
		nil)
	if err != nil {
		c.log.Fatal().Err(err).Msg("Error in binding queue to consume message")
		return err
	}

	deliveries, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	if err != nil {
		c.log.Fatal().Err(err).Msg("Error in consuming message")
		return err
	}

	go func() {
		select {
		case <-c.ctx.Done():
			defer func(ch *amqp.Channel) {
				err := ch.Close()
				if err != nil {
					c.log.Fatal().Err(err).Msgf("failed to close channel closed for for queue: %s", q.Name)
				}
			}(ch)
			c.log.Info().Msgf("channel closed for for queue: %s", q.Name)
			return

		case delivery, ok := <-deliveries:
			{
				if !ok {
					c.log.Fatal().Err(err).Msgf("NOT OK deliveries channel closed for queue: %s", q.Name)
					return
				}

				// Extract headers
				// TODO
				// c.ctx = otel.ExtractAMQPHeaders(c.ctx, delivery.Headers)

				err := c.handler(q.Name, delivery, dependencies)
				if err != nil {
					c.log.Fatal().Err(err).Msg(err.Error())
				}

				consumedMessages = append(consumedMessages, c.queueName)

				h, err := jsoniter.Marshal(delivery.Headers)

				if err != nil {
					c.log.Fatal().Err(err).Msgf("Error in marshalling headers in consumer: %v", string(h))
				}

				err = delivery.Ack(false)
				if err != nil {
					c.log.Fatal().Err(err).Msgf("We didn't get a ack for delivery: %v", string(delivery.Body))
				}
			}
		}
	}()

	c.log.Info().Msgf("Waiting for messages in queue :%s. To exit press CTRL+C", q.Name)

	return nil
}

func (c Consumer[T]) IsConsumed(msg interface{}) bool {
	timeOutTime := 20 * time.Second
	startTime := time.Now()
	timeOutExpired := false
	isConsumed := false

	for {
		if timeOutExpired {
			return false
		}
		if isConsumed {
			return true
		}

		time.Sleep(time.Second * 2)

		typeName := reflect.TypeOf(msg).Name()
		snakeTypeName := strcase.ToSnake(typeName)

		isConsumed = linq.From(consumedMessages).Contains(snakeTypeName)

		timeOutExpired = time.Since(startTime) > timeOutTime
	}
}

func NewConsumer[T any](
	ctx context.Context,
	cfg *RabbitMQConfig,
	conn *amqp.Connection,
	log zerolog.Logger,
	handler func(queue string, msg amqp.Delivery, dependencies T) error,
	exchangeName string,
	exchangeType string,
	queueName string,
	routingKey string,
) IConsumer[T] {
	return &Consumer[T]{
		ctx:          ctx,
		cfg:          cfg,
		conn:         conn,
		log:          log,
		handler:      handler,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queueName:    queueName,
		routingKey:   routingKey,
	}
}

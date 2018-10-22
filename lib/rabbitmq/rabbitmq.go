package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

func New(s string) *RabbitMQ {
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

func (mq *RabbitMQ) Bind(exchange string) {
	err := mq.channel.QueueBind(
		mq.Name,
		"",
		exchange,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	mq.exchange = exchange
}



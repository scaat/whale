package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/gin-gonic/gin/json"
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

func (mq *RabbitMQ) Send(queue string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = mq.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    str,},
	)

	if err != nil {
		panic(err)
	}
}

func (mq *RabbitMQ) Publish(exchange string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = mq.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    str,
		},
	)
	if err != nil {
		panic(err)
	}
}
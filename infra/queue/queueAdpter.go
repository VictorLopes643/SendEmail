package queue

import (
	// "encoding/json"
	// "log"
	// "fmt"
	"github.com/streadway/amqp"
)

type QueueInterface interface {
	Connect() (*amqp.Connection, error)
	DeclareQueue(name string) error
	DeclareExchange(name string) error
	Publish(exchange, queueName, body string) error
	QueueBind(nameQueue, exchange string) error
	Consume(queueName string) (<-chan amqp.Delivery, error)
}

type RabbitMQAdapter struct {
	connection *amqp.Connection
}

func (r *RabbitMQAdapter) Connect() (*amqp.Connection, error) {
	var err error
	r.connection, err = amqp.Dial("amqp://localhost/")
	if err != nil {
		return nil, err
	}
	return r.connection, nil
}

func (r *RabbitMQAdapter) DeclareQueue(name string) error {
	ch, err := r.connection.Channel() // Use the method to get the channel
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQAdapter) DeclareExchange(name string) error {
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		name,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQAdapter) QueueBind(nameQueue, exchange string) error {
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	ch.QueueBind(
		nameQueue, // queue name
		"",        // routing key
		exchange,  // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQAdapter) Publish(exchange, queueName, body string) error {
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQAdapter) Consume(queueName string) (<-chan amqp.Delivery, error) {
	ch, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}

	// Consumir mensagens da fila
	msgs, err := ch.Consume(
		queueName, // Nome da fila
		"",        // Consumer Tag (deixe vazio para gerar um tag único)
		true,      // Auto-Acknowledge (se true, o RabbitMQ considera a mensagem como processada assim que é recebida)
		false,     // Exclusivo (se true, a fila será exclusiva para este consumidor)
		false,     // Não-Local (não será enviado para o consumidor que enviou a mensagem)
		false,     // Atraso (não tem suporte)
		nil,       // Argumentos adicionais (opcional)
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

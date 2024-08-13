package usecase

import (
	"fmt"

	"github.com/victorLopes643/sendEmail-GO/infra/queue"
)

func CreateBindQueue(q queue.QueueInterface, exchangeName, queueName string) error {
	conn, err := q.Connect()
	if err != nil {
		fmt.Println("Erro ao conectar com o RabbitMQ: %v", err)
		return err
	} else {
		fmt.Println("Conectado com sucesso!")
	}
	defer conn.Close()
	err = q.DeclareQueue(queueName)
	if err != nil {
		fmt.Println("Erro ao declarar a fila: %v", err)
		return err
	} else {
		fmt.Println("Fila declarada com sucesso!: %v", queueName)
	}
	err = q.DeclareExchange(exchangeName)
	if err != nil {
		fmt.Println("Erro ao declarar a exchange: %v", err)
		return err
	} else {
		fmt.Println("Exchange declarada com sucesso!: %v", exchangeName)
	}
	err = q.QueueBind(queueName, exchangeName)
	if err != nil {
		fmt.Println("Erro ao associar a fila à exchange: %v", err)
		return err
	} else {
		fmt.Println("Fila associada com sucesso!")
	}
	fmt.Println("Deu boa!")
	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/victorLopes643/sendEmail-GO/infra/email"
	"github.com/victorLopes643/sendEmail-GO/infra/queue"
	"github.com/victorLopes643/sendEmail-GO/internal/usecase"
)

type EmailOptions struct {
	From    string
	To      string
	Subject string
	Body    string
}

func handlerMessage(exchangeName, queueName string) {
	rabbitMQAdapter := &queue.RabbitMQAdapter{}
	_, err := rabbitMQAdapter.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar com o RabbitMQ: %v", err)
	} else {
		log.Println("Conectado com sucesso!")
	}
	for i := 0; i < 10; i++ {
		emailOptions := EmailOptions{
			From:    "victorlopes560@gmail.com",
			To:      "victorlopes560@gmail.com",
			Subject: fmt.Sprintf("Assunto do E-mail %d", i),
			Body:    "Corpo do e-mail",
		}
		body, err := json.Marshal(emailOptions)
		if err != nil {
			log.Printf("Erro ao serializar a mensagem %d para JSON: %v", i, err)
			continue
		}
		err = rabbitMQAdapter.Publish(exchangeName, queueName, string(body))
		if err != nil {
			log.Printf("Erro ao publicar a mensagem %d: %v", i, err)
		} else {
			log.Printf("Mensagem %d publicada com sucesso!", i)
		}
	}
}

func main() {
	rabbitMQAdapter := &queue.RabbitMQAdapter{}
	emailServer := &email.SMTPGmailAdpter{}
	err := usecase.CreateBindQueue(rabbitMQAdapter, "Verus2", "FilaChamado2")
	if err != nil {
		log.Fatalf("Erro ao criar e associar a fila e exchange: %v", err)
	}
	handlerMessage("Verus", "FilaChamado")
	err = usecase.ConsumeQueue(rabbitMQAdapter, emailServer, "FilaChamado")
	if err != nil {
		log.Fatalf("Erro ao consumir a fila: %v", err)
	} else {
		log.Println("Consumindo mensagens...")
	}
}

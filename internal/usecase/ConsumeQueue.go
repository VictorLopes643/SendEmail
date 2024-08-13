package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/victorLopes643/sendEmail-GO/config"
	"github.com/victorLopes643/sendEmail-GO/infra/email"
	"github.com/victorLopes643/sendEmail-GO/infra/queue"
)

func ConsumeQueue(q queue.QueueInterface, es email.SMTPEmailServiceInterface, queueName string) error {
	cfg, _ := config.LoadConfig(".")

	es.ConfigServer(cfg.EmailServer, cfg.EmailPort)
	es.ConfigAuth(cfg.EmailUser, cfg.EmailPass)

	conn, err := q.Connect()
	if err != nil {
		fmt.Println("Erro ao conectar com o RabbitMQ: %v", err)
		return err
	} else {
		fmt.Println("Conectado com sucesso!")
	}
	defer conn.Close()
	msgs, err := q.Consume(queueName)
	if err != nil {
		log.Fatalf("Erro ao consumir a fila: %v", err)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Consumindo mensagens...")
	go func() {
		for msg := range msgs {
			emailOptions := email.EmailOptions{}
			err := json.Unmarshal(msg.Body, &emailOptions)

			if err != nil {
				log.Printf("Erro ao decodificar a mensagem: %v", err)
				continue
			}
			err = es.SendEmail(
				emailOptions.From,
				emailOptions.Subject,
				emailOptions.Body,
			)
			if err != nil {
				log.Printf("Erro ao enviar e-mail: %v", err)
				continue
			}
			fmt.Println("E-mail enviado com sucesso!")
		}
	}()
	<-sigchan
	return nil
}

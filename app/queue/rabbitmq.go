package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang-first-api-rest/database"
	"golang-first-api-rest/models"
	"golang-first-api-rest/util"
	"log"
)

func connectToRabbitMQServer() (*amqp.Channel, *amqp.Connection) {
	// Conecte-se ao servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Erro ao conectar-se ao RabbitMQ: %v", err)
	} else {
		fmt.Println("LOG: Conectado ao RabbitMQ com sucesso")
	}

	// Crie um canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao criar um canal: %v", err)
	}

	return ch, conn
}

func Sender(pessoa *models.Pessoa) {
	ch, conn := connectToRabbitMQServer()

	// Declara uma fila para envio (caso não exista)
	QueueDeclare(ch)

	fmt.Println("LOG: Enviando mensagem para RabbitMQ")
	message := util.Serialize(pessoa)
	err := ch.Publish(
		"",             // Exchange
		"pessoa-queue", // Key
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Erro ao enviar a mensagem: %v", err)
	}

	ch.Close()
	conn.Close()
}

func Consumer() {
	ch, conn := connectToRabbitMQServer()

	// Declara uma fila para envio (caso não exista)
	QueueDeclare(ch)

	// Inscrevendo na pessoa-queue para receber mensagens.
	msgs, err := ch.Consume(
		"pessoa-queue", // Nome da fila
		"",             // Consumer
		true,           // Auto Ack
		false,          // Exclusive
		false,          // No Local
		false,          // No Wait
		nil,            // Args
	)
	if err != nil {
		log.Fatalf("Erro ao criar um consumidor: %v", err)
	}

	fmt.Println("LOG: Esperando por mensagens")

	// Criando um canal para receber mensagens em loop infinito.
	forever := make(chan bool)

	go func() {
		var pessoa models.Pessoa
		for message := range msgs {
			fmt.Println("LOG: Mensagem do RabbitMQ recebida")
			fmt.Println("LOG: processando...")
			util.Deserialize(string(message.Body), &pessoa)
			database.CriaPessoa(&pessoa)
		}
	}()

	<-forever

	ch.Close()
	conn.Close()
}

func QueueDeclare(ch *amqp.Channel) {
	_, err := ch.QueueDeclare(
		"pessoa-queue", // Nome da fila
		true,           // Durable
		false,          // Delete when unused
		false,          // Exclusive
		false,          // No-wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatalf("Erro ao declarar a fila: %v", err)
	}
}

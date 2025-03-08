package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Conn_Rabbit struct {
	Broker *amqp.Connection
	Channel *amqp.Channel
	Err string
}

func GetConnRabbit() *Conn_Rabbit {

	error := ""
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener las variables
	rabbitUser := os.Getenv("R_USER")
	rabbitPassword := os.Getenv("R_PASSWORD")
	rabbitIP := os.Getenv("R_IP")
	rabbitPort := os.Getenv("R_PORT")

	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPassword, rabbitIP, rabbitPort)

    // Conexión a RabbitMQ
    conn, err := amqp.Dial(rabbitUrl)
    if err != nil {
		log.Fatal("Error al abrir una conexión hacia rabbitmq")
    }

    // Abrimos un canal
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal("Error al abrir un canal")
    }

    return &Conn_Rabbit{Broker: conn, Channel: ch, Err: error}

}

func (conn *Conn_Rabbit) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
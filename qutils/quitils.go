package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const SensoryDiscoveryExchange = "SensorDiscovery"
const PersistReadingsQueue = "PersistReading"

const SensorListQueue = "SensorList"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed establish connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		false,
		autoDelete,
		false,
		false,
		nil)

	failOnError(err, "Failed to open a channel")
	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

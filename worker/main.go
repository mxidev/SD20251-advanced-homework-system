package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	// Intentar conectar con RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-container:5672/")
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Intentar abrir un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ No se pudo abrir un canal en RabbitMQ: %v", err)
	}
	defer ch.Close()

	// Declarar la cola
	q, err := ch.QueueDeclare("homeworks", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ No se pudo declarar la cola en RabbitMQ: %v", err)
	}

	// Consumir mensajes de la cola
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ No se pudo consumir mensajes de la cola: %v", err)
	}

	log.Println("📚 Profesor esperando tareas...")

	for msg := range msgs {
		log.Printf("📝 Revisando tarea: %s", msg.Body)
		time.Sleep(2 * time.Second)
		grade := rand.Intn(101) // Nota entre 0 y 100
		log.Printf("📊 Tarea calificada: %s -> Nota %d", msg.Body, grade)
	}
}

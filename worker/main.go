package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("homeworks", false, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	log.Println("📚 Profesor esperando tareas...")

	for msg := range msgs {
		log.Printf("📝 Revisando tarea: %s", msg.Body)
		time.Sleep(2 * time.Second)
		grade := rand.Intn(101) // Nota entre 0 y 100
		log.Printf("📊 Tarea calificada: %s -> Nota %d", msg.Body, grade)
	}
}

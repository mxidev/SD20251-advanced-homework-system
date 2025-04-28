package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "server/proto/homework-system/proto"

	"github.com/streadway/amqp"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHomeworkServiceServer
	ch *amqp.Channel
	q  amqp.Queue
}

func (s *server) SubmitHomework(ctx context.Context, req *pb.HomeworkRequest) (*pb.HomeworkResponse, error) {
	body := req.StudentName + ": " + req.Title
	err := s.ch.Publish("", s.q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Printf("❌ Error al publicar tarea: %v", err)
		return &pb.HomeworkResponse{Status: "Error al enviar tarea"}, err
	}
	log.Printf("📨 Tarea recibida de %s", req.StudentName)
	return &pb.HomeworkResponse{Status: "Tarea enviada correctamente"}, nil
}

func main() {
	var conn *amqp.Connection
	var err error

	// Intentar conectar con RabbitMQ con reintentos
	for i := 0; i < 15; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq-container:5672/")
		if err == nil {
			break
		}
		log.Printf("⚠️ No se pudo conectar a RabbitMQ, reintentando en 2 segundos... (%d/15)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Intentar abrir un canal con reintentos
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ No se pudo abrir un canal en RabbitMQ: %v", err)
	}
	defer ch.Close()

	// Intentar declarar la cola con reintentos
	q, err := ch.QueueDeclare("homeworks", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ No se pudo declarar la cola en RabbitMQ: %v", err)
	}

	// Configurar el servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error al iniciar el listener: %v", err)
	}

	// Crear el servidor gRPC y registrar el servicio
	s := grpc.NewServer()
	pb.RegisterHomeworkServiceServer(s, &server{ch: ch, q: q})
	log.Println("🎓 Servidor gRPC escuchando en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor gRPC: %v", err)
	}
}

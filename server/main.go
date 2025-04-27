package main

import (
	"context"
	"log"
	"net"

	pb "homework-system/proto"

	"github.com/streadway/amqp"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHomeworkServiceServer
	ch amqp.Channel
	q  amqp.Queue
}

func (s *server) SubmitHomework(ctx context.Context, req *pb.HomeworkRequest) (*pb.HomeworkResponse, error) {
	body := req.StudentName + ": " + req.Title
	s.ch.Publish("", s.q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	log.Printf("📨 Tarea recibida de %s", req.StudentName)
	return &pb.HomeworkResponse{Status: "Tarea enviada correctamente"}, nil
}

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("homeworks", false, false, false, false, nil)

	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterHomeworkServiceServer(s, &server{ch: *ch, q: q})
	log.Println("🎓 Servidor gRPC escuchando en :50051")
	s.Serve(lis)
}

package main

import (
	"context"
	"log"
	"time"

	pb "homework-system/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewHomeworkServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := c.SubmitHomework(ctx, &pb.HomeworkRequest{
		StudentName: "Maxi",
		Title:       "Informe de estructuras de datos",
	})

	log.Printf("✅ Servidor respondió: %s", resp.Status)
}

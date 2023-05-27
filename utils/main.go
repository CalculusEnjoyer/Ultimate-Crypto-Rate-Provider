package main

import (
	"context"
	"genesis-test-task/services/email/dispatcher/messages/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:2777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new instance of the gRPC client
	client := proto.NewEmailServiceClient(conn)

	resp, _ := client.SendEmail(context.Background(), &proto.SendEmailRequest{Body: "Check", Subject: "Who", To: "kravchukzxy@gmail.com"})
	print(resp.Error)
}

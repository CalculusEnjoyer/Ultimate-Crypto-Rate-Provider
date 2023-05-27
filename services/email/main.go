package main

import (
	"context"
	"genesis-test-task/services/storage/emails/messages/proto"
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
	client := proto.NewStorageServiceClient(conn)

	response, _ := client.GetAllEmails(context.Background(), &proto.GetAllEmailsRequest{})

	emails := response.Email
	for i := range emails {
		print(emails[i] + " ")
	}
}

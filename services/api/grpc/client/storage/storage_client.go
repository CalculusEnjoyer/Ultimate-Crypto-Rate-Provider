package storage

import (
	"context"
	"genesis-test-task/services/storage/emails/messages/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

const dotEnvPath = "./services/api/.env"

var network string
var port string
var client proto.StorageServiceClient
var connection *grpc.ClientConn

type StorageGRPCClient struct{}

func init() {
	_ = godotenv.Load(dotEnvPath)
	network = os.Getenv("NETWORK")
	port = os.Getenv("STORAGE_SERVICE_PORT")
}

func (c *StorageGRPCClient) AddEmail(request proto.AddEmailRequest) (proto.AddEmailResponse, error) {
	conn := c.getConnection()
	defer conn.Close()

	client = proto.NewStorageServiceClient(conn)

	response, err := client.AddEmail(context.Background(), &request)
	if response == nil {
		return proto.AddEmailResponse{}, err
	}
	return *response, err
}

func (c *StorageGRPCClient) GetAllEmails(request proto.GetAllEmailsRequest) proto.GetAllEmailsResponse {
	conn := c.getConnection()
	defer conn.Close()

	client = proto.NewStorageServiceClient(conn)

	response, err := client.GetAllEmails(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to call GetRate: %v", err)
	}

	return *response
}

func (c *StorageGRPCClient) getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(network+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	return conn
}

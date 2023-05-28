package email

import (
	"context"
	"genesis-test-task/services/email/dispatcher/messages/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

const dotEnvPath = "./services/api/.env"

var network string
var port string
var client proto.EmailServiceClient
var connection *grpc.ClientConn

type EmailGRPCClient struct{}

func init() {
	_ = godotenv.Load(dotEnvPath)
	network = os.Getenv("NETWORK")
	port = os.Getenv("EMAIL_SERVICE_PORT")
}

func (c *EmailGRPCClient) SendEmail(request proto.SendEmailRequest) proto.SendEmailResponse {
	conn := c.getConnection()
	defer conn.Close()

	client = proto.NewEmailServiceClient(conn)

	response, err := client.SendEmail(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to call GetRate: %v", err)
	}

	return *response
}

func (c *EmailGRPCClient) getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(network+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	return conn
}

package currency

import (
	"context"
	"genesis-test-task/services/currency/rate/messages/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

const dotEnvPath = "./services/api/.env"

var network string
var port string
var client proto.RateServiceClient
var connection *grpc.ClientConn

type CurrencyGRPCClient struct{}

func init() {
	_ = godotenv.Load(dotEnvPath)
	network = os.Getenv("NETWORK")
	port = os.Getenv("CURRENCY_SERVICE_PORT")
}

func (c *CurrencyGRPCClient) GetRate(request proto.RateRequest) (proto.RateResponse, error) {
	conn := c.getConnection()
	defer conn.Close()

	client = proto.NewRateServiceClient(conn)

	response, err := client.GetRate(context.Background(), &request)
	if err != nil {
		log.Fatalf("Failed to call GetRate: %v", err)
	}

	return *response, err
}

func (c *CurrencyGRPCClient) getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(network+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	return conn
}

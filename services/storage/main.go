package main

import (
	"genesis-test-task/services/storage/emails"
	"genesis-test-task/services/storage/emails/messages/proto"
	"genesis-test-task/services/storage/emails/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const dotEnvPath = "./services/storage/.env"

func main() {
	run()
}

var StorageFile *os.File

func init() {

}

func run() {
	err := godotenv.Load(dotEnvPath)
	network := os.Getenv("NETWORK")
	port := os.Getenv("PORT")

	service := emails.NewService()
	eps := emails.NewEndpointSet(service)
	grpcServer := transport.NewGRPCServer(eps)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterStorageServiceServer(baseServer, grpcServer)
	lis, err := net.Listen(network, ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

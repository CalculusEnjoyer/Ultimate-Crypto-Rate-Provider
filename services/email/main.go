package main

import (
	"genesis-test-task/services/email/dispatcher"
	"genesis-test-task/services/email/dispatcher/messages/proto"
	"genesis-test-task/services/email/dispatcher/transport"
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

func run() {
	err := godotenv.Load(dotEnvPath)
	network := os.Getenv("NETWORK")
	port := os.Getenv("PORT")

	service := dispatcher.NewService()
	eps := dispatcher.NewEndpointSet(service)
	grpcServer := transport.NewGRPCServer(eps)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterEmailServiceServer(baseServer, grpcServer)
	lis, err := net.Listen(network, ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

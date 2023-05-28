package main

import (
	"genesis-test-task/services/currency/rate"
	"genesis-test-task/services/currency/rate/messages/proto"
	"genesis-test-task/services/currency/rate/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const DotEnvPath = "./services/currency/.env"

func main() {
	run()
}

func run() {
	err := godotenv.Load(DotEnvPath)
	network := os.Getenv("NETWORK")
	port := os.Getenv("PORT")

	service := rate.NewService()
	eps := rate.NewEndpointSet(service)
	grpcServer := transport.NewGRPCServer(eps)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterRateServiceServer(baseServer, grpcServer)
	lis, err := net.Listen(network, ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := baseServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

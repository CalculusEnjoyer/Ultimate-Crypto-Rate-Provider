package main

import (
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"storage.com/storage/emails"
	"storage.com/storage/emails/messages/proto"
	"storage.com/storage/emails/transport"
)

func main() {
	run()
}

var StorageFile *os.File

func init() {

}

func run() {
	err := godotenv.Load()
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

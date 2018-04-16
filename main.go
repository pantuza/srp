package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/pantuza/srp/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":4242"
)


type server struct{}


func (s *server) Ready(ctx context.Context, in *pb.ReadyRequest) (*pb.ReadyResponse, error) {
	return &pb.ReadyResponse{Ready: true}, nil
}


func main() {
	fmt.Println("Service Ready Protocol")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	service := grpc.NewServer()
	pb.RegisterSRPServer(service, &server{})

	reflection.Register(service)
	if err := service.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

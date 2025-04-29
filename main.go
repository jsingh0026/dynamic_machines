package main

import (
	"fmt"
	"log"
	"net"

	pb "machines/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMachineServiceServer(grpcServer, &MachineServiceServer{})

	reflection.Register(grpcServer)

	fmt.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package main

import (
	"log"
	"net"
	"userMocker/core"
	"userMocker/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	port := uint32(lis.Addr().(*net.TCPAddr).Port)

	if err != nil {
		log.Fatalf("Failed to listen to %d %v", port, err)
	}
	s, err := server.InitServer(&lis)
	if err != nil {
		log.Fatalf("Failed to initialize GRPC server %v", err)
	}
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	core.RegisterUserGetterServer(grpcServer, &s)
	core.RegisterUsersGetterServer(grpcServer, &s)
	log.Printf("Initialized GRPC Server, listening on %d ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc %v", err)
	}
}

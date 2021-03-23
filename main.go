package main

import (
	"log"
	"net"
	"userMocker/server"
)

func main() {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen to %d %v", lis.Addr().(*net.TCPAddr).Port, err)
	}
	err = server.InitServer(&lis)
	if err != nil {
		log.Fatalf("Failed to initialize GRPC server %v", err)
	}
}

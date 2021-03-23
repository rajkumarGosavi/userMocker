package client

import (
	"context"
	"log"
	"userMocker/core"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var client core.UserGetterClient

func InitClient(host string) {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatalf("Dial failed %v", err)
	}

	defer conn.Close()

	client = core.NewUserGetterClient(conn)
}

func CallGetUser(data int32) {
	req := core.SingleUserRequest{Id: data}
	res, err := client.GetUser(context.Background(), &req, nil)
	if err != nil {
		log.Fatalf("Unable to get user")
	}
	log.Println(res.GetUser())
}

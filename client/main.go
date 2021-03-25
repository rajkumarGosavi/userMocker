package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"userMocker/core"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var userGetterClient core.UserGetterClient
var usersClient core.UsersGetterClient

func InitClient() error {
	port := uint32(9090)
	opts := setOptions()
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), opts...)
	if err != nil {
		log.Fatalf("Dial failed %v", err)
		return err
	}

	userGetterClient = core.NewUserGetterClient(conn)
	usersClient = core.NewUsersGetterClient(conn)
	return nil
}

func CallGetUser(data uint32) {
	req := core.SingleUserRequest{Id: data}
	res, err := userGetterClient.GetUser(context.Background(), &req, nil)
	if err != nil {
		log.Fatalf("Unable to get user")
	}
	log.Println(res.GetUser())
}

func setOptions() (opts []grpc.DialOption) {
	opts = append(opts, grpc.WithInsecure())
	return
}

func main() {
	err := InitClient()
	if err != nil {
		log.Fatalf("Failed to initialize clients %v", err)
	}
	var users, user string
	flag.StringVar(&users, "uids", "0", "-ids <USER_ID1>,<USER_ID2>,<USER_ID3>..... \n 0 is not a valid Id")
	flag.StringVar(&user, "uid", "0", "-user <USER_ID>\n 0 is not a valid Id")

	flag.Parse()
	switch {
	case user != "0":
		i, err := strconv.ParseInt(user, 10, 64)
		if err != nil {
			log.Fatalf("Invalid Input\n %v\n Value Should range between %d to %d", err, 0, math.MaxUint32)
		}
		resp, err := userGetterClient.GetUser(context.Background(), &core.SingleUserRequest{Id: uint32(i)})
		if err != nil {
			log.Fatalf("Failed to Fetch user %v", err)
		}
		log.Println(resp.GetUser())
	case users != "0":
		ids := strings.Split(users, ",")
		finalIds := []uint32{}
		for _, id := range ids {
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Fatalf("Invalid Input %v\n Value Should range between %d to %d", err, 0, math.MaxUint32)
			}
			finalIds = append(finalIds, uint32(i))
			resp, err := usersClient.GetUsers(context.Background(), &core.UsersRequest{Id: finalIds})
			if err != nil {
				log.Fatalf("Failed to Fetch Users %v", err)
			}
			log.Println(resp.GetUser())
		}
	default:
		flag.Usage()
	}
	defer conn.Close()
}

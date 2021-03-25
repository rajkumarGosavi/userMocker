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

// var conn *grpc.ClientConn
var userGetterClient core.UserGetterClient
var usersClient core.UsersGetterClient

// InitClient - initialize dial grpc connection and registers the  grpc clients
func InitClient() (conn *grpc.ClientConn, err error) {
	port := uint32(9090)
	opts := setOptions()
	conn, err = grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), opts...)
	if err != nil {
		log.Fatalf("Dial failed %v", err)
		return nil, err
	}

	userGetterClient = core.NewUserGetterClient(conn)
	usersClient = core.NewUsersGetterClient(conn)
	return
}

// setOptions - used to set "grpc.DialOption"s
func setOptions() (opts []grpc.DialOption) {
	opts = append(opts, grpc.WithInsecure())
	return
}

func main() {

	conn, err := InitClient()
	if err != nil {
		log.Fatalf("Failed to initialize clients %v", err)
	}

	if conn == nil {
		log.Fatalf("Connection Failed to established")
	}
	defer conn.Close()

	var users, user string
	flag.StringVar(&users, "uids", "0", "-ids <USER_ID1>,<USER_ID2>,<USER_ID3>..... \n 0 is not a valid Id")
	flag.StringVar(&user, "uid", "0", "-user <USER_ID>\n 0 is not a valid Id")

	flag.Parse()
	for {
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
			log.Println()
			fmt.Println(resp.GetUser())
			fmt.Println()
			user = "0"
		case users != "0":
			ids := strings.Split(users, ",")
			finalIds := []uint32{}
			for _, id := range ids {
				i, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					log.Fatalf("Invalid Input %v\n Value Should range between %d to %d", err, 0, math.MaxUint32)
				}
				finalIds = append(finalIds, uint32(i))
			}
			resp, err := usersClient.GetUsers(context.Background(), &core.UsersRequest{Id: finalIds})
			if err != nil {
				log.Fatalf("Failed to Fetch Users %v", err)
			}

			log.Printf("\nUsers Details:\n")
			for _, u := range resp.GetUser() {
				fmt.Println(u)
			}
			fmt.Println()
			users = "0"
		default:
			return
		}
	}
}

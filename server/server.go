package server

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"
	"userMocker/core"
)

// Server - defines the server type
type Server struct {
	users []*core.User
	Port  uint32
	core.UnimplementedUserGetterServer
	core.UnimplementedUsersGetterServer
}

// GetUser - Will get a particular user from the users list according to their id
func (s *Server) GetUser(ctx context.Context, request *core.SingleUserRequest) (*core.SingleUserResponse, error) {
	u := &core.User{}
	for _, usr := range s.users {
		if usr.GetId() == request.GetId() {
			u = usr
			return &core.SingleUserResponse{User: u}, nil
		}
	}
	return &core.SingleUserResponse{User: &core.User{}}, nil
}

// GetUsers - Will get users from the users list according to their ids
func (s *Server) GetUsers(ctx context.Context, request *core.UsersRequest) (*core.UsersResponse, error) {
	u := make([]*core.User, 0)
	for _, id := range request.Id {
		for _, usr := range s.users {
			if usr.Id == id {
				u = append(u, usr)
			}
		}
	}
	return &core.UsersResponse{User: u}, nil
}

// InitServer - initializes a server for grpc and adds users data to server itself
func InitServer(lis *net.Listener) (Server, error) {
	port := uint32((*lis).Addr().(*net.TCPAddr).Port)
	temp := []*core.User{}
	rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})
	log.Println("Printing user ids:")
	for _, t := range temp {
		log.Println(t.GetId())
	}
	s := Server{users: temp, Port: port}

	return s, nil
}

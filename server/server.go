package server

import (
	"context"
	"log"
	"math/rand"
	"net"
	"userMocker/core"

	"google.golang.org/grpc"
)

type Server struct {
	users []*core.User
	core.UnimplementedUserGetterServer
}

func (s *Server) GetUser(ctx context.Context, request *core.SingleUserRequest) (*core.Response, error) {
	u := &core.User{}
	for _, usr := range s.users {
		if usr.GetId() == request.GetId() {
			u = usr
			return &core.Response{User: []*core.User{u}}, nil
		}
	}
	return &core.Response{User: []*core.User{}}, nil
}
func (s *Server) GetUsers(ctx context.Context, request *core.UsersRequest) (*core.Response, error) {
	u := make([]*core.User, 0)
	for _, id := range request.Id {
		for _, usr := range s.users {
			if usr.Id == id {
				u = append(u, usr)
			}
		}
	}
	return &core.Response{User: u}, nil
}

func InitServer(lis *net.Listener) error {
	port := (*lis).Addr().(*net.TCPAddr).Port
	temp := []*core.User{}
	// rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})

	s := Server{users: temp}
	grpcServer := grpc.NewServer()

	core.RegisterUserGetterServer(grpcServer, &s)
	log.Printf("Initialized GRPC Server, listening on %d ", port)
	if err := grpcServer.Serve(*lis); err != nil {
		log.Fatalf("Failed to serve grpc %v", err)
		return err
	}
	return nil
}

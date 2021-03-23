package server

import (
	"context"
	"log"
	"math/rand"
	"net"
	"testing"
	"userMocker/core"
)

func TestGetUser(t *testing.T) {
	s := Server{}
	temp := []*core.User{}
	// rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})
	ids := []int32{}
	for _, v := range temp {
		ids = append(ids, v.GetId())
	}
	s.users = temp

	tests := []struct {
		id   int32
		want int
		name string
	}{
		{
			id:   1,
			want: 0,
			name: "Wrong ID",
		},
		{
			id:   ids[0],
			want: 1,
			name: "Correct ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := core.SingleUserRequest{Id: tt.id}
			resp, err := s.GetUser(context.Background(), &req)
			if err != nil {
				t.Errorf("Unexpected Error %v for %d", err, tt.id)
			}
			val := resp.GetUser()
			got := len(val)
			if got != tt.want {
				t.Errorf("got %d want %d for %d %v", got, tt.want, tt.id, val)
				for _, v := range val {
					log.Println("**", v)
				}
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	s := Server{}
	temp := []*core.User{}
	// rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Int31(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})
	ids := []int32{}
	for _, v := range temp {
		ids = append(ids, v.GetId())
	}
	s.users = temp

	tests := []struct {
		ids  []int32
		want int
		name string
	}{
		{
			ids:  []int32{},
			want: 0,
			name: "Zero ids",
		},
		{
			ids:  ids[0:1],
			want: 1,
			name: "Single id",
		},
		{
			ids:  ids[0:2],
			want: 2,
			name: "Single id",
		},
		{
			ids:  ids[:],
			want: len(ids),
			name: "All ids",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := core.UsersRequest{Id: tt.ids}
			resp, err := s.GetUsers(context.Background(), &req)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			got := len(resp.GetUser())
			if got != tt.want {
				t.Errorf("got %d want %d for %v", got, tt.want, req.String())
			}
		})
	}
}

func TestInitServer(t *testing.T) {

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}
	port := lis.Addr().(*net.TCPAddr).Port
	got := InitServer(&lis)
	var want error

	if got != want {
		t.Errorf("got %v want %v for port %d", got, want, port)
	}

	// nettest.TestConn(t, func() (c1 net.Conn, c2 net.Conn, stop func(), err error) {
	// 	c1, c2 = net.Pipe()
	// 	// if err = InitServer(""); err != nil {
	// 	// 	t.Errorf("Connection establishment failed on %s %v", c1.LocalAddr().String(), err)
	// 	// }
	// 	stop = func() {
	// 		c1.Close()
	// 		c2.Close()
	// 	}
	// 	return
	// })
}

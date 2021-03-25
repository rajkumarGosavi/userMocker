package server

import (
	"context"
	"math/rand"
	"testing"
	"userMocker/core"
)

func TestGetUser(t *testing.T) {
	s := Server{}
	temp := []*core.User{}
	// rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})
	ids := []uint32{}
	for _, v := range temp {
		ids = append(ids, v.GetId())
	}
	s.users = temp

	tests := []struct {
		id   uint32
		want uint32
		name string
	}{
		{
			id:   1,
			want: 0,
			name: "Wrong ID",
		},
		{
			id:   ids[0],
			want: temp[0].Id,
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
			got := val
			if got.Id != tt.want {
				t.Errorf("got %v want %v for %d %v", got, tt.want, tt.id, val)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	s := Server{}
	temp := []*core.User{}
	// rand.Seed(time.Now().UnixNano())
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User1", City: "City1", Height: 5.4, Married: false, Phone: 1998800123})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User2", City: "City2", Height: 4.4, Married: true, Phone: 9999999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User3", City: "City3", Height: 6.4, Married: true, Phone: 9123999999})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User4", City: "City4", Height: 5.1, Married: false, Phone: 9999998909})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User5", City: "City5", Height: 5.9, Married: true, Phone: 9999456788})
	temp = append(temp, &core.User{Id: rand.Uint32(), Fname: "User6", City: "City6", Height: 5.11, Married: false, Phone: 9918273645})
	ids := []uint32{}
	for _, v := range temp {
		ids = append(ids, v.GetId())
	}
	s.users = temp

	tests := []struct {
		ids  []uint32
		want int
		name string
	}{
		{
			ids:  []uint32{},
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

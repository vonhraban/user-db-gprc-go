package main

import (
	"fmt"
	"log"
	"net"

	"context"

	"google.golang.org/grpc"

	pb "github.com/vonhraban/user-db-grpc-go/user"
)

const ( 
	port = ":50051" 
)

// server is used to implement user.UserServer.
type server struct {
	savedUsers []*pb.UserEntity
}

func (s *server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := request.GetUser()
	log.Printf("Saving new user: %+v", user)

	s.savedUsers = append(s.savedUsers, user)

	log.Print("Saved")


	return &pb.CreateUserResponse{Id: user.Id, Success: true}, nil
}

func (s *server) GetUserById(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.GetUserResponse, error) {
	log.Printf("Looking for user: %+v", request.Id)

	for i := range s.savedUsers {
		if s.savedUsers[i].Id == request.Id {
			log.Printf("Found, returning %v", s.savedUsers[i])

			return &pb.GetUserResponse{
				User:  s.savedUsers[i],
			}, nil
		}
	}

	log.Printf("Did not find user %d", request.Id)

	return &pb.GetUserResponse{
		Error: &pb.Error{
			Message: fmt.Sprintf("Could not find user with id %d", request.Id),
		},
	}, nil
}


func main() { 
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
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

func (s server) CreateUser(ctx context.Context, userRequest *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := userRequest.GetUser()
	s.savedUsers = append(s.savedUsers, user)

	return &pb.CreateUserResponse{Id: user.Id, Success: true}, nil
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

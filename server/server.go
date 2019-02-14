package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vonhraban/user-db-grpc-go/user"
)

// server is used to implement user.UserServer.
type server struct {
	repo userRepository
}

func newServer(repo userRepository) *server {
	return &server{repo: repo}
}

func (s *server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := request.GetUser()
	log.Printf("Saving new user: %+v", user)

	if err := s.repo.Add(user); err != nil {
		return nil, err
	}

	log.Print("Saved")

	return &pb.CreateUserResponse{Id: user.Id, Success: true}, nil
}

func (s *server) GetUserByID(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.GetUserResponse, error) {
	log.Printf("Looking for user: %+v", request.Id)

	user, err := s.repo.GetById(request.Id)
	if err != nil {
		return nil, err
	}

	if user != nil {
		log.Printf("Found user %v", user)

		return &pb.GetUserResponse{
			User: user,
		}, nil
	}

	log.Printf("Did not find user %d", request.Id)

	return &pb.GetUserResponse{
		Error: &pb.Error{
			Message: fmt.Sprintf("Could not find user with id %d", request.Id),
		},
	}, nil
}

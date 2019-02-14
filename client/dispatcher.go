package main

import (
	"context"
	"fmt"
	pb "github.com/vonhraban/user-db-grpc-go/user"
)


type dispatcher struct {
	client pb.UserClient
}

func newDispatcher(client pb.UserClient) dispatcher {
	return dispatcher{client: client}
}

func (d *dispatcher) addUser(user *pb.UserEntity) (int, error) {
	request := &pb.CreateUserRequest{
		User: user,
	}

	response, err := d.client.CreateUser(context.Background(), request)

	if err != nil {
		return 0, err
	}

	if !response.Success {
		return 0, fmt.Errorf("Something went wrong creating the user", response)
	}

	return int(response.Id), nil
}

func (d *dispatcher) getUser(id int) (*pb.UserEntity, error) {
	request := &pb.GetUserByIdRequest{
		Id: int32(id),
	}

	response, err := d.client.GetUserByID(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return response.User, nil
}



package main

import (
	pb "github.com/vonhraban/user-db-grpc-go/user"
)


type userRepository interface {
	Add(user *pb.UserEntity) error
	GetById(id int32) (*pb.UserEntity, error)
}
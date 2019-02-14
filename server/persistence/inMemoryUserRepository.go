package persistence

import (
	pb "github.com/vonhraban/user-db-grpc-go/user"
)


type InMemoryUserRepository struct{
	users []*pb.UserEntity
}

func (r *InMemoryUserRepository) Add(user *pb.UserEntity) error {
	r.users = append(r.users, user)

	return nil
}

func (r *InMemoryUserRepository) GetById(id int32) (*pb.UserEntity, error) {
	for i := range r.users {
		if r.users[i].Id == id {
			return r.users[i], nil
		}
	}
	return nil, nil
}
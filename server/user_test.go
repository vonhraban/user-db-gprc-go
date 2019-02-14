package main

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/vonhraban/user-db-grpc-go/server/persistence"
	pb "github.com/vonhraban/user-db-grpc-go/user"
	"reflect"
)

type journey struct {
	server *server
}

var userTemplate = pb.UserEntity{
	Name: "Jane Smith",
	Emails: []*pb.UserEntity_Email{
		{
			Address:"jane.smith@gmail.com",
			IsPrimary: true,
		},
		{
			Address:"jane154@aol.com",
			IsPrimary: false,
		},
	},
	PhoneNumbers: []*pb.UserEntity_Phone{
		{
			Number: "555-454-453",
			Type: pb.UserEntity_Phone_LANDLINE,
		},
		{
			Number: "123-456-789",
			Type: pb.UserEntity_Phone_MOBILE,
		},
	},
}

func (j *journey) userWithIdDoesNotExist(arg1 int) error {
	// do nothing here
	return nil
}

func (j *journey) iAddAUserWithId(id int32) error {
	userToPersist := userTemplate
	userToPersist.Id = id

	request := &pb.CreateUserRequest{
		User: &userToPersist,
	}

	if _, err := j.server.CreateUser(context.Background(), request); err != nil {
		return err
	}

	return nil
}

func (j *journey) theUserWithIdHasToExist(id int32) error {
	retrievedUser, err := j.server.repo.GetById(id)
	expectedUser := userTemplate
	expectedUser.Id = id

	if err != nil {
		panic(err)
	}

	if retrievedUser == nil {
		return fmt.Errorf(
			"could not find user with id %d",
			id,
		)
	}

	if !reflect.DeepEqual(expectedUser, *retrievedUser) {
		return fmt.Errorf(
			"expected user: %+v. Actual: %+v",
			expectedUser,
			retrievedUser,
		)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	userRepository := persistence.InMemoryUserRepository{}
	server := newServer(&userRepository)


	journey := &journey{
		server: server,
	}

	s.Step(`^user with id (\d+) does not exist$`, journey.userWithIdDoesNotExist)
	s.Step(`^I add a user with id (\d+)$`, journey.iAddAUserWithId)
	s.Step(`^the user with id (\d+) has to exist$`, journey.theUserWithIdHasToExist)
}

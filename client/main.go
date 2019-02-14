package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/vonhraban/user-db-grpc-go/user"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to grpc server: %v", err)
	}

	defer func() {
		err := conn.Close()

		if err != nil {
			log.Fatalf("Error closing connection %v", err)
		}
	}()

	// Creates a new CustomerClient
	client := pb.NewUserClient(conn)
	dispatcher := newDispatcher(client)

	user := &pb.UserEntity{
		Id: 43443,
		Name: "Jane Smith",
		Emails: []*pb.UserEntity_Email{
			{
				Address:"jane.smith@gmail.com",
				IsPrimary: true,
			},
			{
				Address:"jane154@aol.com",
				IsPrimary: true,
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

	id, err := dispatcher.addUser(user)

	if err != nil {
		log.Fatalf("could not persist user: %v", err)
	}

	fmt.Printf("Added new user id %d \n", id)

	userId := 456789
	user, err = dispatcher.getUser(userId)

	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Printf("Could not find a user with id %d \n", userId)
	} else {
		fmt.Printf("Got user with id %d: %v \n", userId, user)
	}



	userId = 43443
	user, err = dispatcher.getUser(userId)

	if err != nil {
		panic(err)
	}

	if user == nil {
		fmt.Printf("Could not find a user with id %d \n", userId)
	} else {
		fmt.Printf("Got user with id %d: %v \n", userId, user)
	}
}
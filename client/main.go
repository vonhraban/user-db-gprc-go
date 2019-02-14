package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/vonhraban/user-db-grpc-go/user"
)

const (
	address = "localhost:50051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createUser(client pb.UserClient, request *pb.CreateUserRequest) {
	response, err := client.CreateUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	if response.Success {
		log.Printf("A new user has been added with id: %d", response.Id)
	}
}

func getUserById(client pb.UserClient, request *pb.GetUserByIdRequest) {
	response, err := client.GetUserById(context.Background(), request)
	if err != nil {
		log.Fatalf("An error has occured: %v", err)
	}

	if response.Error != nil {
		log.Printf("Server returned an error: %v", response.Error)
	}

	if response.User != nil {
		log.Printf("Fetched user: %v", response.User)
	}
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to grpc server: %v", err)
	}

	defer func() {
		err := conn.Close()

		if err != nil {
			panic("error closing connection")
		}
	}()

	// Creates a new CustomerClient
	client := pb.NewUserClient(conn)

	user := &pb.CreateUserRequest{
		User: &pb.UserEntity{
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
		},
	}

	// Create a new customer
	createUser(client, user)

	getUserByIdRequest := &pb.GetUserByIdRequest{
		Id: 4344343,
	}
	getUserById(client, getUserByIdRequest)

	getUserByIdRequest = &pb.GetUserByIdRequest{
		Id: 43443,
	}
	getUserById(client, getUserByIdRequest)
}
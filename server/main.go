package main

import (
	"github.com/vonhraban/user-db-grpc-go/server/persistence"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/vonhraban/user-db-grpc-go/user"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	userRepository := persistence.InMemoryUserRepository{}
	pb.RegisterUserServer(s, newServer(&userRepository))

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

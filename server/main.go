package main

import (
	"context"
	"gorpc/contract"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf(err.Error())
	}

	serve := grpc.NewServer()
	contract.RegisterHelloServiceServer(serve, &server{})

	log.Println("Server is running......")

	serve.Serve(listen)
}

func (s *server) Hello(ctx context.Context, in *contract.Request) (*contract.Person, error) {
	return &contract.Person{
		Id:         "1",
		FirstName:  "Awal",
		MiddleName: "Tengah",
		LastName:   "Akhir",
		Email:      "dynastymasra@gmail.com",
		Phone:      "0987654321",
		Gender:     contract.Gender_MALE,
	}, nil
}

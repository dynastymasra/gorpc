package main

import (
	"context"
	"gorpc/contract"
	"log"
	"net"

	"strings"

	"google.golang.org/grpc"
)

type server struct {
	persons []*contract.Person
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf(err.Error())
	}

	serve := grpc.NewServer()
	contract.RegisterPersonServiceServer(serve, &server{})

	log.Println("Server is running......")

	serve.Serve(listen)
}

func (s *server) CreatePerson(ctx context.Context, person *contract.Person) (*contract.Person, error) {
	s.persons = append(s.persons, person)
	return person, nil
}

func (s *server) FilterPerson(filter *contract.Filter, stream contract.PersonService_FilterPersonServer) error {
	for _, person := range s.persons {
		if !strings.Contains(person.Name, filter.Name) {
			continue
		}

		if err := stream.Send(person); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetAllPerson(empty *contract.Empty, stream contract.PersonService_GetAllPersonServer) error {
	for _, person := range s.persons {
		if err := stream.Send(person); err != nil {
			return err
		}
	}

	return nil
}

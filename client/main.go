package main

import (
	"log"

	"context"
	"gorpc/contract"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}

	client := contract.NewHelloServiceClient(conn)
	person, err := client.Hello(context.Background(), &contract.Request{"1"})
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(person)
}

package main

import (
	"log"

	"context"
	"gorpc/contract"

	"io"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()

	person := &contract.Person{
		Id:   "1",
		Name: "Number One",
		Email: []*contract.Email{{
			Email:     "email_1",
			IsPrimary: true,
		}, {
			Email:     "email_2",
			IsPrimary: false,
		},
		},
		Phone: "098765431",
		Address: &contract.Address{
			State:      "state",
			Street:     "street",
			City:       "city",
			PostalCode: "postal code",
		},
		Gender: contract.Gender_MALE,
	}

	client := contract.NewPersonServiceClient(conn)
	res, err := client.CreatePerson(context.Background(), person)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(res)

	person = &contract.Person{
		Id:   "2",
		Name: "Number Two",
		Email: []*contract.Email{{
			Email:     "email_1",
			IsPrimary: true,
		}, {
			Email:     "email_2",
			IsPrimary: false,
		},
		},
		Phone: "098765431",
		Address: &contract.Address{
			State:      "state",
			Street:     "street",
			City:       "city",
			PostalCode: "postal code",
		},
		Gender: contract.Gender_MALE,
	}

	res, err = client.CreatePerson(context.Background(), person)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(res)

	person = &contract.Person{
		Id:   "3",
		Name: "Number Three",
		Email: []*contract.Email{{
			Email:     "email_1",
			IsPrimary: true,
		}, {
			Email:     "email_2",
			IsPrimary: false,
		},
		},
		Phone: "098765431",
		Address: &contract.Address{
			State:      "state",
			Street:     "street",
			City:       "city",
			PostalCode: "postal code",
		},
		Gender: contract.Gender_FEMALE,
	}

	res, err = client.CreatePerson(context.Background(), person)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(res)

	all, err := client.GetAllPerson(context.Background(), &contract.Empty{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	filter, err := client.FilterPerson(context.Background(), &contract.Filter{Name: "Three"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	for {
		one, err := all.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
		}
		log.Printf("All Person = %v \n", one)
	}

	for {
		two, err := filter.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		log.Printf("Filter = %v \n", two)
	}
}

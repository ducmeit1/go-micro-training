package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-training/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewFPTPeopleClient(conn)

	people, err := client.CreatePeople(context.Background(), &pb.People{
		Name:    "Client Duc",
		Address: "Client FPT Software",
		Age:     25,
	})
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

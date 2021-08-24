package main

import (
	"fmt"
	"gin-training/grpc/people-grpc/handlers"
	"gin-training/grpc/people-grpc/repositories"
	"gin-training/helper"
	"gin-training/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	peopleRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewPeopleHandler(peopleRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFPTPeopleServer(s, h)

	fmt.Println("Listen at port: 2222")

	s.Serve(listen)
}

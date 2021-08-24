package main

import (
	"fmt"
	"gin-training/grpc/people-job-grpc/handlers"
	"gin-training/grpc/people-job-grpc/repositories"
	"gin-training/helper"
	"gin-training/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	peopleConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	jobConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	peopleClient := pb.NewFPTPeopleClient(peopleConn)
	jobClient := pb.NewFPTJobClient(jobConn)

	err = helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":2224")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	peopleJobRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewPeopleJobHandler(peopleClient, jobClient, peopleJobRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFPTPeopleJobServer(s, h)

	fmt.Println("Listen at port: 2224")

	s.Serve(listen)
}

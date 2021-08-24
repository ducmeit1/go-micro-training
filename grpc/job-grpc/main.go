package main

import (
	"fmt"
	"gin-training/grpc/job-grpc/handlers"
	"gin-training/grpc/job-grpc/repositories"
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

	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	jobRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewJobHandler(jobRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFPTJobServer(s, h)

	fmt.Println("Listen at port: 2223")

	s.Serve(listen)
}

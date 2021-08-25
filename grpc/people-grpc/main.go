package main

import (
	"fmt"
	"gin-training/grpc/people-grpc/handlers"
	"gin-training/grpc/people-grpc/repositories"
	"gin-training/helper"
	"gin-training/intercepter"
	"gin-training/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
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

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
	)

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

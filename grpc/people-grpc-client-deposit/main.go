package main

import (
	"context"
	"fmt"
	"gin-training/pb"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	peopleConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	peopleClient := pb.NewFPTPeopleClient(peopleConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := peopleClient.DepositAccountBalance(ctx)
	if err != nil {
		panic(err)
	}

	done := make(chan bool, 1)

	go func() {
		for i := 0; i < 10; i++ {
			err := stream.Send(&pb.ChangeAccountBalanceRequest{
				PeopleId:      "11",
				BalanceChange: float64(i * 100),
			})
			if err != nil {
				panic(err)
			}
		}
		done <- true
	}()

	<-done

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Printf("PeopleID: %v, Balance remain: %v, Update At: %v\n", res.PeopleId, res.BlanaceRemain, res.UpdatedAt.AsTime().String())

	<-ctx.Done()
}

package main

import (
	"context"
	"fmt"
	"gin-training/pb"
	"io"
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

	stream, err := peopleClient.ChangeAccountBalance(ctx)
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			err := stream.Send(&pb.ChangeAccountBalanceRequest{
				PeopleId:      "11",
				BalanceChange: float64(i * 100),
			})
			if err != nil {
				panic(err)
			}
			time.Sleep(300 * time.Millisecond)
		}
		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}

			fmt.Printf("PeopleID: %v, Balance Change:%v Balance remain: %v, Update At: %v\n", res.PeopleId, res.BalanceChange, res.BlanaceRemain, res.UpdatedAt.AsTime().String())
		}
	}()

	<-ctx.Done()
}

package main

import (
	"gin-training/api/people-api/handlers"
	"gin-training/pb"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	peopleConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	peopleClient := pb.NewFPTPeopleClient(peopleConn)

	//Handler for GIN Gonic
	h := handlers.NewPeopleHandler(peopleClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()

	//Create routes
	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreatePeople)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}

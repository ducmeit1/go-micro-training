package main

import (
	"gin-training/api/people-api/handlers"
	"gin-training/middleware"
	"gin-training/pb"
	custom_validator "gin-training/validator"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := handlers.NewPeopleHandler(peopleClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("minage", custom_validator.ValidAgeValidator(int64(18)))
	}

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/create", h.CreatePeople)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}

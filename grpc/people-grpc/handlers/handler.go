package handlers

import (
	"context"
	"gin-training/grpc/people-grpc/models"
	"gin-training/grpc/people-grpc/repositories"
	"gin-training/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PeopleHandler struct {
	pb.UnimplementedFPTPeopleServer
	peopleRepository repositories.PeopleRepository
}

func NewPeopleHandler(peopleRepository repositories.PeopleRepository) (*PeopleHandler, error) {
	return &PeopleHandler{
		peopleRepository: peopleRepository,
	}, nil
}

func (h *PeopleHandler) CreatePeople(ctx context.Context, in *pb.People) (*pb.People, error) {
	people, err := h.peopleRepository.CreatePeople(ctx, &models.People{
		Name:    in.Name,
		Age:     in.Age,
		Address: in.Address,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.People{
		Id:       int64(people.ID),
		Name:     people.Name,
		Age:      people.Age,
		Address:  people.Address,
		Contacts: []*pb.Contact{},
	}, nil
}

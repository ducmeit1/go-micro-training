package handlers

import (
	"context"
	"database/sql"
	"gin-training/grpc/people-grpc/models"
	"gin-training/grpc/people-grpc/repositories"
	"gin-training/pb"

	"github.com/google/uuid"
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
	pRequest := &models.People{
		ID:   uuid.New(),
		Name: in.Name,
		Age:  in.Age,
		Slut: in.Slut,
		Address: sql.NullString{
			String: in.Address,
			Valid:  true,
		},
		Contact: make([]*models.Contact, 0),
	}

	for _, v := range in.Contacts {
		pRequest.Contact = append(pRequest.Contact, &models.Contact{
			ID:          uuid.New(),
			PeopleID:    pRequest.ID,
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	people, err := h.peopleRepository.CreatePeople(ctx, pRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pResponse := &pb.People{
		Id:       people.ID.String(),
		Slut:     people.Slut,
		Name:     people.Name,
		Age:      people.Age,
		Address:  people.Address.String,
		Contacts: make([]*pb.Contact, 0),
	}

	for _, v := range people.Contact {
		pResponse.Contacts = append(pResponse.Contacts, &pb.Contact{
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	return pResponse, nil
}

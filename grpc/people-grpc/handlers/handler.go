package handlers

import (
	"context"
	"database/sql"
	"gin-training/grpc/people-grpc/models"
	"gin-training/grpc/people-grpc/repositories"
	"gin-training/grpc/people-grpc/requests"
	"gin-training/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func findContactIDIndex(contactID uuid.UUID, in []*models.Contact) int {
	for i := range in {
		if in[i].Id == contactID {
			return i
		}
	}
	return -1
}

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
		Id:   uuid.New(),
		Name: in.Name,
		Age:  in.Age,
		Slut: in.Slut,
		Address: sql.NullString{
			String: in.Address,
			Valid:  true,
		},
		Contacts: make([]*models.Contact, 0),
	}

	for _, v := range in.Contacts {
		pRequest.Contacts = append(pRequest.Contacts, &models.Contact{
			Id:          uuid.New(),
			PeopleID:    pRequest.Id,
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
		Id:       people.Id.String(),
		Slut:     people.Slut,
		Name:     people.Name,
		Age:      people.Age,
		Address:  people.Address.String,
		Contacts: make([]*pb.Contact, 0),
	}

	for _, v := range people.Contacts {
		pResponse.Contacts = append(pResponse.Contacts, &pb.Contact{
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	return pResponse, nil
}

func (h *PeopleHandler) UpdatePeople(ctx context.Context, in *pb.People) (*pb.People, error) {
	people, err := h.peopleRepository.GetPeopleByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Address != "" {
		people.Address = sql.NullString{
			Valid:  true,
			String: in.Address,
		}
	}

	if in.Age != 0 && in.Age > 18 {
		people.Age = in.Age
	}

	if in.Name != "" {
		people.Name = in.Name
	}

	if in.Slut != "" {
		people.Slut = in.Slut
	}

	if len(in.Contacts) > 0 {
		for _, v := range in.Contacts {
			index := findContactIDIndex(uuid.MustParse(v.Id), people.Contacts)
			if index < 0 {
				return nil, status.Errorf(codes.NotFound, "contact id %v not found to update", v.Id)
			}

			if v.Email != "" {
				people.Contacts[index].Email = v.Email
			}

			if v.Fax != "" {
				people.Contacts[index].Fax = v.Fax
			}

			if v.PhoneNumber != "" {
				people.Contacts[index].PhoneNumber = v.PhoneNumber
			}
		}
	}

	newPeople, err := h.peopleRepository.UpdatePeople(ctx, people)
	if err != nil {
		return nil, err
	}

	pResponse := &pb.People{
		Id:       newPeople.Id.String(),
		Slut:     newPeople.Slut,
		Name:     newPeople.Name,
		Age:      newPeople.Age,
		Address:  newPeople.Address.String,
		Contacts: make([]*pb.Contact, 0),
	}

	for _, v := range newPeople.Contacts {
		pResponse.Contacts = append(pResponse.Contacts, &pb.Contact{
			Id:          v.Id.String(),
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	return pResponse, nil
}

func (h *PeopleHandler) FindPeople(ctx context.Context, in *pb.FindPeopleRequest) (*pb.People, error) {
	var (
		people = &models.People{}
		err    error
	)

	if in.Id == "" && in.Slut == "" {
		return nil, status.Error(codes.InvalidArgument, "id or slut is required")
	}

	if in.Id != "" {
		people, err = h.peopleRepository.GetPeopleByID(ctx, uuid.MustParse(in.Id))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	if in.Slut != "" {
		people, err = h.peopleRepository.GetPeopleBySlut(ctx, in.Slut)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}

	pRes := &pb.People{
		Id:       people.Id.String(),
		Slut:     people.Slut,
		Name:     people.Name,
		Age:      people.Age,
		Address:  people.Address.String,
		Contacts: []*pb.Contact{},
	}

	for _, v := range people.Contacts {
		pRes.Contacts = append(pRes.Contacts, &pb.Contact{
			Id:          v.Id.String(),
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	return pRes, nil
}

func (h *PeopleHandler) ListPeoples(ctx context.Context, in *pb.ListPeopleRequest) (*pb.ListPeopleResponse, error) {
	peoples, err := h.peopleRepository.ListPeoples(ctx, &requests.ListPeopleRequest{
		Age: in.Age,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.ListPeopleResponse{
		Peoples: []*pb.People{},
	}

	err = copier.CopyWithOption(&pRes.Peoples, &peoples, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *PeopleHandler) DeletePeople(ctx context.Context, in *pb.DeletePeopleRequest) (*pb.Empty, error) {
	if err := h.peopleRepository.DeletePeople(ctx, uuid.MustParse(in.Id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return &pb.Empty{}, nil
}

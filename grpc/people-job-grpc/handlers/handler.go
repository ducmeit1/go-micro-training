package handlers

import (
	"context"
	"gin-training/grpc/people-job-grpc/models"
	"gin-training/grpc/people-job-grpc/repositories"
	"gin-training/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PeopleJobHandler struct {
	peopleClient        pb.FPTPeopleClient
	jobClient           pb.FPTJobClient
	peopleJobRepository repositories.PeopleJobRepositories
	pb.UnimplementedFPTPeopleJobServer
}

func NewPeopleJobHandler(peopleClient pb.FPTPeopleClient,
	jobClient pb.FPTJobClient,
	peopleJobRepository repositories.PeopleJobRepositories) (*PeopleJobHandler, error) {
	return &PeopleJobHandler{
		peopleClient:        peopleClient,
		jobClient:           jobClient,
		peopleJobRepository: peopleJobRepository,
	}, nil
}

func (h *PeopleJobHandler) AssignPeopleJob(ctx context.Context, in *pb.AssignPeopleJobRequest) (*pb.PeopleJob, error) {
	if in.PeopleId == "" {
		return nil, status.Error(codes.InvalidArgument, "people_id is required")
	}

	if in.JobId == "" {
		return nil, status.Error(codes.InvalidArgument, "job_id is required")
	}

	people, err := h.peopleClient.FindPeople(ctx, &pb.FindPeopleRequest{
		Id: in.PeopleId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "people_id is not found")
			}
		} else {
			return nil, err
		}
	}

	job, err := h.jobClient.FindJob(ctx, &pb.FindJobRequest{
		Id: in.JobId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "job_id is not found")
			}
		} else {
			return nil, err
		}
	}

	peopleJob := &models.PeopleJob{
		Id:         uuid.New(),
		PeopleID:   uuid.MustParse(people.Id),
		PeopleSlut: people.Slut,
		PeopleName: people.Name,
		JobID:      uuid.MustParse(job.Id),
		JobName:    job.Name,
		JobRanking: in.JobRanking,
		Salary:     in.Salary,
		Active:     in.Active,
	}

	res, err := h.peopleJobRepository.AssignPeopleJob(ctx, peopleJob)
	if err != nil {
		return nil, err
	}

	pRes := &pb.PeopleJob{}
	err = copier.Copy(&pRes, &res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *PeopleJobHandler) ListPeopleJob(ctx context.Context, in *pb.ListPeopleJobRequest) (*pb.ListPeopleJobResponse, error) {
	panic("not implemented") // TODO: Implement
}

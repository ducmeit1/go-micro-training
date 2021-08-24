package handlers

import (
	"context"
	"database/sql"
	"gin-training/grpc/job-grpc/models"
	"gin-training/grpc/job-grpc/repositories"
	"gin-training/grpc/job-grpc/requests"
	"gin-training/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JobHandler struct {
	pb.UnimplementedFPTJobServer
	jobRepository repositories.JobRepositories
}

func NewJobHandler(jobRepository repositories.JobRepositories) (*JobHandler, error) {
	return &JobHandler{
		jobRepository: jobRepository,
	}, nil
}

func (h *JobHandler) CreateJob(ctx context.Context, req *pb.Job) (*pb.Job, error) {
	job := &models.Job{}

	err := copier.Copy(&job, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.jobRepository.CreateJob(ctx, job)
	if err != nil {
		return nil, err
	}

	pRes := &pb.Job{}
	err = copier.Copy(&pRes, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *JobHandler) FindJob(ctx context.Context, req *pb.FindJobRequest) (*pb.Job, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	res, err := h.jobRepository.FindJob(ctx, uuid.MustParse(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "job is not found")
		}
		return nil, err
	}

	pRes := &pb.Job{}
	err = copier.Copy(&pRes, &res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

func (h *JobHandler) ListJob(ctx context.Context, req *pb.ListJobRequest) (*pb.ListJobResponse, error) {
	listJob := &requests.ListJobRequest{}

	err := copier.Copy(&listJob, &req)
	if err != nil {
		return nil, err
	}

	res, err := h.jobRepository.ListJob(ctx, listJob)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "job not found")
		}
		return nil, err
	}

	pRes := &pb.ListJobResponse{}
	err = copier.Copy(&pRes.Jobs, res)
	if err != nil {
		return nil, err
	}

	return pRes, nil
}

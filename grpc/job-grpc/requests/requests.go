package requests

import "github.com/google/uuid"

type ListJobRequest struct {
	Level int64
}

type FindJobRequest struct {
	Id uuid.UUID
}

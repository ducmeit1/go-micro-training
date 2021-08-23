package requests

import "github.com/gofrs/uuid"

type UpdatePeopleRequest struct {
	Id      uuid.UUID
	Name    string
	Age     int64
	Address string
	Slut    string
}

type ListPeopleRequest struct {
	Age int64
}

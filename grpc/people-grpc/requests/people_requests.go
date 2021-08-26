package requests

import "github.com/google/uuid"

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

type UpdateAccountBalanceRequest struct {
	PeopleId       uuid.UUID
	AccountBalance float64
}

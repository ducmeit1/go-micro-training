package requests

import "github.com/gofrs/uuid"

type UpdatePeopleRequest struct {
	ID      uuid.UUID
	Name    string
	Age     int64
	Address string
	Slut    string
}

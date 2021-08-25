package requests

type CreatePeopleRequest struct {
	Name     string                  `json:"name" binding:"required"`
	Age      int64                   `json:"age" binding:"required,minage"`
	Address  string                  `json:"address" binding:"max=256,min=6"`
	Slut     string                  `json:"slut" binding:"required,min=6,max=256"`
	Contacts []*CreateContactRequest `json:"contacts,omitempty"`
}

type CreateContactRequest struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Fax         string `json:"fax"`
}

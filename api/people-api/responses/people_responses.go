package responses

type PeopleResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Age      int64            `json:"age"`
	Slut     string           `json:"slut"`
	Contacts []*PeopleContact `json:"contacts"`
}

type PeopleContact struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Fax         string `json:"fax"`
}

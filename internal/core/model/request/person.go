package request

type AddPersonRequest struct {
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Address   string `json:"Address"`
	City      string `json:"City"`
}

type DeletePersonRequest struct {
	PersonsID string `json:"PersonsID"`
}

type GetPersonsByFirstName struct {
	FirstName string `json:"firstname"`
}

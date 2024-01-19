package response

type AddUserDataResponse struct {
	DisplayName string `json:"displayName"`
}

type PersonsDataResponse struct {
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Address   string `json:"Address"`
	City      string `json:"City"`
}

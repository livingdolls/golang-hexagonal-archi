package request

type AddTodoRequest struct {
	PersonID string `json:"personid"`
	Title    string `json:"title"`
}

type DeleteTodoRequest struct {
	Id string `json:"id"`
}

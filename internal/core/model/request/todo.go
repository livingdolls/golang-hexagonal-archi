package request

type AddTodoRequest struct {
	PersonID string `json:"personid"`
	Title    string `json:"title"`
}

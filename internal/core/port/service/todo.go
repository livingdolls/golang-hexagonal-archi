package service

import (
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
)

type TodoService interface {
	AddTodo(request *request.AddTodoRequest) *response.HttpResponse
	GetListTodo() *response.HttpResponse
}

package service

import (
	"gotest/internal/core/dto"
	"gotest/internal/core/entity/error_code"
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
	"gotest/internal/core/port/repository"
	"gotest/internal/core/port/service"
	"log"
	"time"
)

var (
	invalidPerson = "invalid person"
	invalidTitle  = "invalid title"
)

type todoService struct {
	todoRepo repository.TodoRepository
}

// DeleteTodoById implements service.TodoService.
func (t *todoService) DeleteTodoById(request *request.DeleteTodoRequest) *response.HttpResponse {
	err := t.todoRepo.Delete(request.Id)

	if err != nil {
		log.Println(err.Error())

		if err == repository.TodoNotFound {
			return t.errorResponse(error_code.NotFound, error_code.InvalidNotFoundMsg)
		}

		return t.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	return t.successResponse(request.Id)
}

func (t *todoService) GetListTodo() *response.HttpResponse {
	res, err := t.todoRepo.GetList()

	if err != nil {
		log.Println(err.Error())

		return t.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	return t.successResponse(res)
}

// AddTodo implements service.TodoService.
func (t *todoService) AddTodo(request *request.AddTodoRequest) *response.HttpResponse {
	if len(request.PersonID) == 0 {
		return t.errorResponse(error_code.InvalidRequest, invalidPerson)
	}

	if len(request.Title) == 0 {
		return t.errorResponse(error_code.InvalidRequest, invalidTitle)
	}

	todoDTO := dto.TodoDTO{
		PersonID:  request.PersonID,
		Title:     request.Title,
		CreatedAt: time.Now(),
	}

	err := t.todoRepo.Insert(todoDTO)

	if err != nil {
		if err == repository.DuplicateTodo {
			return t.errorResponse(error_code.DuplicateTodo, err.Error())
		}
		log.Println(err.Error())
		return t.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	todoResponse := response.TodoDataResponse{
		PersonID:  request.PersonID,
		Title:     request.Title,
		CreatedAt: time.Now(),
	}

	return t.successResponse(todoResponse)

}

func NewTodoService(todoRepo repository.TodoRepository) service.TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t todoService) successResponse(data any) *response.HttpResponse {
	return &response.HttpResponse{
		Data:         data,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}

func (t todoService) errorResponse(code error_code.ErrorCode, message string) *response.HttpResponse {
	return &response.HttpResponse{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

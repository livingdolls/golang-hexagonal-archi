package service

import (
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
)

type PersonService interface {
	AddPerson(request *request.AddPersonRequest) *response.HttpResponse
	GetListPersons() *response.HttpResponse
	DeletePerson(request *request.DeletePersonRequest) *response.HttpResponse
	GetPersonByFirstName(request *request.GetPersonsByFirstName) *response.HttpResponse
}

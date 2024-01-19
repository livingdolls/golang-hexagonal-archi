package service

import (
	"gotest/internal/core/dto"
	"gotest/internal/core/entity/error_code"
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
	"gotest/internal/core/port/repository"
	"gotest/internal/core/port/service"
	"log"
)

const (
	invalidName = "invalid name"
)

type personService struct {
	personRepo repository.PersonRepository
}

// DeletePerson implements service.PersonService.
func (p personService) DeletePerson(request *request.DeletePersonRequest) *response.HttpResponse {
	if len(request.PersonsID) == 0 {
		return p.errorResponse(error_code.InvalidRequest, error_code.InvalidRequestErrMsg)
	}

	err := p.personRepo.DeletePerson(request.PersonsID)

	if err != nil {
		log.Panicln(err.Error())
		return p.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	return p.successResponse(request.PersonsID)
}

// GetListPersons implements service.PersonService.
func (p personService) GetListPersons() *response.HttpResponse {
	res, err := p.personRepo.GetPersons()

	if err != nil {
		log.Println(err.Error())
		return p.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	return p.successResponse(res)
}

// AddPerson implements service.PersonService.
func (p personService) AddPerson(request *request.AddPersonRequest) *response.HttpResponse {
	if len(request.FirstName) == 0 {
		return p.errorResponse(error_code.InvalidRequest, invalidName)
	}

	personDTO := dto.PersonDTO{
		LastName:  request.LastName,
		FirstName: request.FirstName,
		Address:   request.Address,
		City:      request.City,
	}

	err := p.personRepo.Insert(personDTO)

	if err != nil {
		if err == repository.DuplicatePerson {
			return p.errorResponse(error_code.DuplicateUser, err.Error())
		}

		log.Println(err.Error())
		return p.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// create response

	addPerson := response.AddUserDataResponse{
		DisplayName: personDTO.FirstName,
	}

	return p.successResponse(addPerson)
}

func NewPersonService(personRepo repository.PersonRepository) service.PersonService {
	return &personService{
		personRepo: personRepo,
	}
}

func (p personService) successResponse(data any) *response.HttpResponse {
	return &response.HttpResponse{
		Data:         data,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}

func (p personService) errorResponse(code error_code.ErrorCode, message string) *response.HttpResponse {
	return &response.HttpResponse{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

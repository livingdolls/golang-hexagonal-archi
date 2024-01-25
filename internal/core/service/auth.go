package service

import (
	"gotest/internal/core/entity/error_code"
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
	"gotest/internal/core/port/repository"
	"gotest/internal/core/port/service"
)

type AuthService struct {
	ts   service.TokenService
	repo repository.PersonRepository
}

func NewAuthService(ts service.TokenService, repo repository.PersonRepository) *AuthService {
	return &AuthService{
		repo: repo,
		ts:   ts,
	}
}

func (as AuthService) Login(request *request.GetPersonsByFirstName) *response.HttpResponse {
	if len(request.FirstName) == 0 {
		return as.errorResponse(error_code.InvalidRequest, error_code.InvalidRequestErrMsg)
	}

	res, err := as.repo.GetPersonsByFirstName(request.FirstName)

	if err != nil {
		if err == repository.NoPersonsInRow {
			return as.errorResponse(error_code.NotFound, error_code.InvalidNotFoundMsg)
		}

		return as.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	accessToken, err := as.ts.CreateToken(res)

	if err != nil {
		if len(accessToken) == 0 {
			return as.errorResponse(error_code.InternalError, err.Error())
		}
		return as.errorResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	return as.successResponse(accessToken)
}

func (p AuthService) successResponse(data any) *response.HttpResponse {
	return &response.HttpResponse{
		Data:         data,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
		Status:       true,
	}
}

func (p AuthService) errorResponse(code error_code.ErrorCode, message string) *response.HttpResponse {
	return &response.HttpResponse{
		ErrorCode:    code,
		ErrorMessage: message,
		Status:       false,
	}
}

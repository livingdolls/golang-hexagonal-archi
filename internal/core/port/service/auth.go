package service

import (
	"gotest/internal/core/dto"
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
)

type TokenService interface {
	CreateToken(person *dto.PersonDTO) (string, error)
	ValidateToken(token string) (*dto.TokenPayload, error)
}

type AuthService interface {
	Login(request *request.GetPersonsByFirstName) *response.HttpResponse
}

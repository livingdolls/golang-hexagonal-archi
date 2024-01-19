package repository

import (
	"errors"
	"gotest/internal/core/dto"
)

var DuplicatePerson = errors.New("duplicate person")

type PersonRepository interface {
	Insert(person dto.PersonDTO) error
	GetPersons() ([]dto.PersonDTO, error)
	DeletePerson(id string) error
}

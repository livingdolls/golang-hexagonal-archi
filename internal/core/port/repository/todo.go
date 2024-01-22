package repository

import (
	"errors"
	"gotest/internal/core/dto"
)

var DuplicateTodo = errors.New("duplicate todo")

type TodoRepository interface {
	Insert(todo dto.TodoDTO) error
	GetList() ([]dto.TodoDTO, error)
}

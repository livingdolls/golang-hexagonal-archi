package repository

import (
	"errors"
	"gotest/internal/core/dto"
)

var (
	DuplicateTodo = errors.New("duplicate todo")
	TodoNotFound  = errors.New("todo not found")
)

type TodoRepository interface {
	Insert(todo dto.TodoDTO) error
	GetList() ([]dto.TodoDTO, error)
	Delete(id string) error
}

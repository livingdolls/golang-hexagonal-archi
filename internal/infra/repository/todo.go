package repository

import (
	"context"
	"errors"
	"gotest/internal/core/dto"
	"gotest/internal/core/port/repository"
	"log"
	"strconv"
	"strings"
)

type todoRepository struct {
	db repository.Database
}

var (
	todoInsertErr = errors.New("failed insert todo")
)

const (
	insertTodo = "INSERT INTO todo ( " +
		"`personId`," +
		"`title`," +
		"`createdAt`" +
		") VALUES (?,?,?)"

	getTodo    = "SELECT * FROM todo"
	deleteTodo = "DELETE FROM todo WHERE id = ?"
)

// Delete implements repository.TodoRepository.
func (t *todoRepository) Delete(id string) error {
	i, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	res, err := t.db.GetDB().ExecContext(context.Background(), deleteTodo, i)

	if err != nil {
		log.Println(err.Error())

		return err
	}

	numbRow, errRow := res.RowsAffected()

	if errRow != nil {
		return errRow
	}

	if numbRow == 0 {
		return repository.TodoNotFound
	}

	return nil
}

func (t *todoRepository) GetList() ([]dto.TodoDTO, error) {
	var todo dto.TodoDTO
	var todos []dto.TodoDTO

	res, err := t.db.GetDB().Query(getTodo)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(
			&todo.PersonID,
			&todo.Id,
			&todo.Title,
			&todo.CreatedAt,
			&todo.DoneAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

// Insert implements repository.TodoRepository.
func (t *todoRepository) Insert(todo dto.TodoDTO) error {
	res, err := t.db.GetDB().Exec(
		insertTodo,
		todo.PersonID,
		todo.Title,
		todo.CreatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), duplicateEntryMsg) {
			return repository.DuplicateTodo
		}

		return err
	}

	numbRow, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if numbRow != numberRowInserted {
		return todoInsertErr
	}

	return nil
}

func NewTodoRepository(db repository.Database) repository.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

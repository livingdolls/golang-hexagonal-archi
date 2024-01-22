package dto

import (
	"database/sql"
	"time"
)

type PersonDTO struct {
	PersonsID string
	LastName  string
	FirstName string
	Address   string
	City      string
}

type TodoDTO struct {
	PersonID  string
	Id        int
	Title     string
	CreatedAt time.Time
	DoneAt    sql.NullTime
}

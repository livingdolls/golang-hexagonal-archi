package repository

import (
	"errors"
	"gotest/internal/core/dto"
	"gotest/internal/core/port/repository"
	"strings"
)

const (
	duplicateEntryMsg = "Duplicate entry"
	numberRowInserted = 1
)

var (
	instertUserErr = errors.New("failed to insert user")
)

const (
	insertStatement = "INSERT INTO Persons ( " +
		"`PersonID`," +
		"`LastName`," +
		"`FirstName`," +
		"`Address`," +
		"`City`" +
		") VALUES (?,?,?,?,?)"
)

type personRepository struct {
	db repository.Database
}

// DeletePerson implements repository.PersonRepository.
func (p personRepository) DeletePerson(id string) error {
	_, err := p.db.GetDB().Query("DELETE FROM Persons WHERE PersonID = ?", id)

	if err != nil {
		return err
	}

	return nil
}

// GetPersons implements repository.PersonRepository.
func (p personRepository) GetPersons() ([]dto.PersonDTO, error) {
	var persons dto.PersonDTO
	var listpersons []dto.PersonDTO

	res, err := p.db.GetDB().Query("SELECT * FROM Persons")

	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(
			&persons.PersonsID,
			&persons.LastName,
			&persons.FirstName,
			&persons.Address,
			&persons.City,
		)

		if err != nil {
			return nil, err
		}

		listpersons = append(listpersons, persons)
	}

	return listpersons, nil
}

// Insert implements repository.PersonRepository.
func (p personRepository) Insert(person dto.PersonDTO) error {
	res, err := p.db.GetDB().Exec(
		insertStatement,
		1,
		person.LastName,
		person.FirstName,
		person.Address,
		person.City,
	)

	if err != nil {
		if strings.Contains(err.Error(), duplicateEntryMsg) {
			return repository.DuplicatePerson
		}

		return err
	}

	numbRow, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if numbRow != numberRowInserted {
		return instertUserErr
	}

	return nil
}

func NewPersonRepository(db repository.Database) repository.PersonRepository {
	return &personRepository{
		db: db,
	}
}

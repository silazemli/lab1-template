package server

import (
	"github.com/silazemli/lab1-template/internal/server/models"
)

type personStorage interface {
	CreatePerson(person models.Person) (models.Person, error)
	GetPerson(id string) (models.Person, error)
	GetAll() ([]models.Person, error)
	DeletePerson(id string) error
	PatchPerson(id string, person models.Person) error
}

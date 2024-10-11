package storage

import (
	"os"

	"github.com/silazemli/lab1-template/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewDB() (*storage, error) {
	dsn := os.Getenv("PGDSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &storage{}, err
	}
	return &storage{db}, nil
}

func (stg *storage) CreatePerson(person models.Person) (models.Person, error) {
	err := stg.db.Table("persons").Create(&person)
	if err.Error != nil {
		return models.Person{}, err.Error
	}
	return person, nil
}

func (stg *storage) GetPerson(id string) (models.Person, error) {
	person := models.Person{}
	err := stg.db.Table("persons").Where("id = ?", id).Take(&person).Error
	if err != nil {
		return models.Person{}, err
	}
	return person, nil
}

func (stg *storage) GetAll() ([]models.Person, error) {
	persons := []models.Person{}
	err := stg.db.Table("persons").Find(&persons).Error
	if err != nil {
		return []models.Person{}, err
	}
	return persons, nil
}

func (stg *storage) DeletePerson(id string) error {
	err := stg.db.Table("persons").Where("id = ?", id).Delete(models.Person{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (stg *storage) PatchPerson(id string, person models.Person) error {
	err := stg.db.Table("persons").Where("id = ?", id).Updates(&person).Error
	if err != nil {
		return err
	}
	return nil
}

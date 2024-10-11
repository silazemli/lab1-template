package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/silazemli/lab1-template/internal/server/models"
)

type personStorageTest struct {
}

func (pst *personStorageTest) CreatePerson(person models.Person) (models.Person, error) {
	return models.Person{}, nil
}

func (pst *personStorageTest) GetPerson(id string) (models.Person, error) {
	return models.Person{}, nil
}

func (pst *personStorageTest) GetAll() ([]models.Person, error) {
	return nil, nil
}

func (pst *personStorageTest) DeletePerson(id string) error {
	return nil
}

func (pst *personStorageTest) PatchPerson(id string, person models.Person) error {
	return nil
}

func TestServer_createPerson(t *testing.T) {
	srv := NewServer(&personStorageTest{})
	r := httptest.NewRequest(http.MethodPost, "/test", nil)
	w := httptest.NewRecorder()
	ctx := srv.srv.NewContext(r, w)
	err := srv.CreatePerson(ctx)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	code := w.Result().StatusCode
	if code != 201 {
		t.Errorf("Test failed")
	}
}

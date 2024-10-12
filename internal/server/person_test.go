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

func TestServer_CreatePerson(t *testing.T) {
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

func TestServer_PatchPerson(t *testing.T) {
	srv := NewServer(&personStorageTest{})
	r := httptest.NewRequest(http.MethodPatch, "/test", nil)
	w := httptest.NewRecorder()
	ctx := srv.srv.NewContext(r, w)
	err := srv.PatchPerson(ctx)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	code := w.Result().StatusCode
	if code != 200 {
		t.Errorf("Test Failed")
	}
}

func TestServer_GetPerson(t *testing.T) {
	srv := NewServer(&personStorageTest{})
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	ctx := srv.srv.NewContext(r, w)
	err := srv.GetPerson(ctx)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	code := w.Result().StatusCode
	if code != 200 {
		t.Errorf("Test Failed")
	}
}

func TestServer_GetAll(t *testing.T) {
	srv := NewServer(&personStorageTest{})
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	ctx := srv.srv.NewContext(r, w)
	err := srv.GetAll(ctx)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	code := w.Result().StatusCode
	if code != 200 {
		t.Errorf("Test Failed")
	}
}

func TestServer_DeleteAll(t *testing.T) {
	srv := NewServer(&personStorageTest{})
	r := httptest.NewRequest(http.MethodDelete, "/test", nil)
	w := httptest.NewRecorder()
	ctx := srv.srv.NewContext(r, w)
	err := srv.DeletePerson(ctx)
	if err != nil {
		t.Errorf("Should not produce an error")
	}
	code := w.Result().StatusCode
	if code != 204 {
		t.Errorf("Test Failed")
	}
}

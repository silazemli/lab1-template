package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/silazemli/lab1-template/internal/server/models"
)

type server struct {
	srv echo.Echo
	db  personStorage
}

func NewServer(db personStorage) server {
	srv := server{}
	srv.db = db
	srv.srv = *echo.New()
	srv.srv.GET("/persons/:id", srv.GetPerson)
	srv.srv.GET("/persons", srv.GetAll)
	srv.srv.POST("/persons", srv.CreatePerson)
	srv.srv.PATCH("/persons/:id", srv.PatchPerson)
	srv.srv.DELETE("/persons/:id", srv.DeletePerson)

	return srv
}

func (srv *server) Start() error {
	err := srv.srv.Start(":8080")
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) GetPerson(ctx echo.Context) error {
	person, err := srv.db.GetPerson(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusFound, person)
}

func (srv *server) GetAll(ctx echo.Context) error {
	persons, err := srv.db.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusFound, persons)

}

func (srv *server) CreatePerson(ctx echo.Context) error {
	person := models.Person{}
	err := ctx.Bind(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	person, err = srv.db.CreatePerson(person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusCreated, person)
}

func (srv *server) DeletePerson(ctx echo.Context) error {
	id := ctx.Param("id")
	err := srv.db.DeletePerson(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusNoContent, models.Person{})
}

func (srv *server) PatchPerson(ctx echo.Context) error {
	person := models.Person{}
	err := ctx.Bind(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	err = srv.db.PatchPerson(ctx.Param("id"), person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return ctx.JSON(http.StatusOK, person)
}

package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id    string `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func GetUserList(c echo.Context) error {
	u := &User{
		Id:    "1",
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func GetUserByID(c echo.Context) error {
	u := &User{
		Id:    c.Param("id"),
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

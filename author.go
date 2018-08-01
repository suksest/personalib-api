package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getAllAuthor(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	authors := []Author{}
	db.Find(&authors)
	return c.JSON(http.StatusOK, authors)
}

func getAuthorByID(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	author := new(Author)

	id, _ := strconv.Atoi(c.Param("id"))

	db.First(&author, id)
	return c.JSON(http.StatusOK, author)
}

func createAuthor(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	author := new(Author)

	if err := c.Bind(author); err != nil {
		return err
	}

	db.NewRecord(author)
	db.Create(&author)

	return c.JSON(http.StatusOK, author)
}

func updateAuthor(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	b := new(Author)
	author := new(Author)

	if err := c.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&author, id).Updates(b)

	return c.JSON(http.StatusOK, author)
}

func deleteAuthor(c echo.Context) error {
	db := OpenDB()
	defer db.Close()

	author := new(Author)

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&author, id)
	db.Delete(&author)

	return c.NoContent(http.StatusNoContent)
}

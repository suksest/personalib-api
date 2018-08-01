package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getAllBook(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	books := []Book{}
	db.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func getBookByID(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	book := new(Book)

	id, _ := strconv.Atoi(c.Param("id"))

	db.First(&book, id)
	return c.JSON(http.StatusOK, book)
}

func createBook(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	book := new(Book)

	if err := c.Bind(book); err != nil {
		return err
	}

	db.NewRecord(book)
	db.Create(&book)

	return c.JSON(http.StatusOK, book)
}

func updateBook(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	b := new(Book)
	book := new(Book)

	if err := c.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&book, id).Updates(b)

	return c.JSON(http.StatusOK, book)
}

func deleteBook(c echo.Context) error {
	db := OpenDB()
	defer db.Close()

	book := new(Book)

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&book, id)
	db.Delete(&book)

	return c.NoContent(http.StatusNoContent)
}

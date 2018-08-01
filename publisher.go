package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getAllPublisher(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	publishers := []Publisher{}
	db.Model(&publishers).Related(&Book{})
	return c.JSON(http.StatusOK, publishers)
}

func getPublisherByID(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	publisher := new(Publisher)

	id, _ := strconv.Atoi(c.Param("id"))

	db.First(&publisher, id)
	return c.JSON(http.StatusOK, publisher)
}

func createPublisher(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	publisher := new(Publisher)

	if err := c.Bind(publisher); err != nil {
		return err
	}

	db.NewRecord(publisher)
	db.Create(&publisher)

	return c.JSON(http.StatusOK, publisher)
}

func updatePublisher(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	b := new(Publisher)
	publisher := new(Publisher)

	if err := c.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&publisher, id).Updates(b)

	return c.JSON(http.StatusOK, publisher)
}

func deletePublisher(c echo.Context) error {
	db := OpenDB()
	defer db.Close()

	publisher := new(Publisher)

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&publisher, id)
	db.Delete(&publisher)

	return c.NoContent(http.StatusNoContent)
}

package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getAllCategory(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	categories := []Category{}
	db.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func getCategoryByID(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	category := new(Category)

	id, _ := strconv.Atoi(c.Param("id"))

	db.First(&category, id)
	return c.JSON(http.StatusOK, category)
}

func createCategory(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	category := new(Category)

	if err := c.Bind(category); err != nil {
		return err
	}

	db.NewRecord(category)
	db.Create(&category)

	return c.JSON(http.StatusOK, category)
}

func updateCategory(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	b := new(Category)
	category := new(Category)

	if err := c.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&category, id).Updates(b)

	return c.JSON(http.StatusOK, category)
}

func deleteCategory(c echo.Context) error {
	db := OpenDB()
	defer db.Close()

	category := new(Category)

	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&category, id)
	db.Delete(&category)

	return c.NoContent(http.StatusNoContent)
}

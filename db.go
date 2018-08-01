package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

//OpenDB function
func OpenDB() *gorm.DB {
	db, err := gorm.Open("postgres",
		"host=localhost port=5432 user=sukma dbname=personalib password=openpgpwd")
	if err != nil {
		panic(err)
	}

	return db
}

//Migration function
func Migration() {
	db := OpenDB()
	defer db.Close()
	db.AutoMigrate(&Category{}, &Author{}, &Publisher{}, &Book{}, &BookCategory{})
	db.Model(&Book{}).AddForeignKey("author_id", "authors(id)", "RESTRICT", "RESTRICT")
	db.Model(&Book{}).AddForeignKey("publisher_id", "publishers(id)", "RESTRICT", "RESTRICT")
	db.Model(&BookCategory{}).AddForeignKey("book_id", "books(id)", "RESTRICT", "RESTRICT")
	db.Model(&BookCategory{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
}

func index(c echo.Context) error {
	message := new(Message)
	message.Message = "success"
	return c.JSON(http.StatusOK, message)
}

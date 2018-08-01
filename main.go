package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Migration()
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Routers
	e.GET("/", index)
	e.POST("/book", createBook)
	e.GET("/book/:id", getBookByID)
	e.GET("/book", getAllBook)
	e.PUT("/book/:id", updateBook)
	e.DELETE("/book/:id", deleteBook)
	e.POST("/author", createAuthor)
	e.GET("/author/:id", getAuthorByID)
	e.GET("/author", getAllAuthor)
	e.PUT("/author/:id", updateAuthor)
	e.DELETE("/author/:id", deleteAuthor)
	e.POST("/publisher", createPublisher)
	e.GET("/publisher/:id", getPublisherByID)
	e.GET("/publisher", getAllPublisher)
	e.PUT("/publisher/:id", updatePublisher)
	e.DELETE("/publisher/:id", deletePublisher)
	e.POST("/category", createCategory)
	e.GET("/category/:id", getCategoryByID)
	e.GET("/category", getAllCategory)
	e.PUT("/category/:id", updateCategory)
	e.DELETE("/category/:id", deleteCategory)

	//Start Server
	e.Logger.Fatal(e.Start(":1323"))

}

package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Connection went wrong: %s", err.Error()))
	}

	repo := repository.NewBookRepository(db)
	service := service.NewBookService(repo)
	controller := controller.NewBookController(service)

	r.GET("/books", controller.ListBooks)
	r.GET("/book/:id", controller.GetBookById)
	r.POST("/book", controller.CreateBook)
	r.PUT("/book/:id", controller.UpdateBookById)
	r.DELETE("/book/:id", controller.DeleteBookById)

	r.Run(":5000")
}

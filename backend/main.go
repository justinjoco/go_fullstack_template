package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	connStr := "postgres://admin:password@postgres/app_db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Connection went wrong: %s", err.Error()))
	}
	defer db.Close()

	maxAttempts := 3
	for i := 0; i < maxAttempts; i++ {
		err := db.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	repo := repository.NewBookRepository(db)
	svc := service.NewBookService(repo)
	ctr := controller.NewBookController(svc)

	r.GET("/books", ctr.ListBooks)
	r.GET("/book/:id", ctr.GetBookById)
	r.POST("/book", ctr.CreateBook)
	r.PUT("/book/:id", ctr.UpdateBookById)
	r.DELETE("/book/:id", ctr.DeleteBookById)

	r.Run(":5000")
}

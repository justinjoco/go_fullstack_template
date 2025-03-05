package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func connectToDB() (*sql.DB, error) {

	connStr := "postgres://admin:password@postgres/app_db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("DB connection failed")
		return nil, err
	}

	err = db.Ping()
	if err == nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	r := gin.Default()

	var db *sql.DB
	var err error

	for {
		db, err = connectToDB()
		if err == nil {
			log.Println("Connection success")
			break
		}
		log.Println("Waiting to connect again")
		time.Sleep(1 * time.Second)
	}

	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	healthCheckController := controller.NewHealthCheckController()

	r.GET("/books", bookController.ListBooks)
	r.GET("/book/:id", bookController.GetBookById)
	r.POST("/book", bookController.CreateBook)
	r.PUT("/book/:id", bookController.UpdateBookById)
	r.DELETE("/book/:id", bookController.DeleteBookById)

	r.GET("/health_check", healthCheckController.HealthCheck)

	_ = r.Run(":5000")
}

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
		log.Fatal("DB connection failed")
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

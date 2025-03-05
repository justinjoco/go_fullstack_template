package main

import (
	"app/controller"
	"app/repository"
	"app/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {

	dsn := "postgres://admin:password@postgres/app_db"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("DB connection failed")
		return nil, err
	}

	return db, nil
}

func connectToCache() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		DB:       0,
		Password: "mypassword",
		Username: "default",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Redis connection failed")
		return nil, err
	}
	return rdb, nil
}
func main() {
	r := gin.Default()

	db, err := connectToDB()
	if err != nil {
		log.Fatal("DB connect failure")
	}
	sqlDB, _ := db.DB()
	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
	}
	log.Println("DB connection success")
	defer sqlDB.Close()

	rdb, err := connectToCache()
	if err != nil {
		log.Fatal("Redis connect failure")
	}
	log.Println("Redis connection success")

	bookRepo := repository.NewBookRepository(db, rdb)
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

package main

import (
	"app/controller"
	"app/models"
	"app/repository"
	"app/service"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
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

func connectToCache(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		DB:       0,
		Password: "mypassword",
		Username: "default",
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Redis connection failed")
		return nil, err
	}
	return rdb, nil
}

func setupCache(ctx context.Context, db *gorm.DB, rdb *redis.Client) {
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		log.Fatal("unable to seed redis")
	}
	// Iterate over users and store them in Redis
	for _, book := range books {
		// Convert user to JSON
		bookData, err := json.Marshal(book)
		if err != nil {
			log.Printf("Error marshalling book %d: %v", book.Id, err)
			continue
		}

		// Store user in Redis (using the user ID as the key)
		key := "book:" + book.Id.String()
		err = rdb.Set(ctx, key, bookData, time.Hour).Err()
		if err != nil {
			log.Printf("Error storing book %d in Redis: %v", book.Id, err)
		} else {
			log.Printf("Book %d cached in Redis", book.Id)
		}
	}
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
	ctx := context.Background()

	rdb, err := connectToCache(ctx)
	if err != nil {
		log.Fatal("Redis connect failure")
	}
	log.Println("Redis connection success")

	setupCache(ctx, db, rdb)
	bookRepo := repository.NewBookRepository(db, rdb)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	healthCheckController := controller.NewHealthCheckController()

	corsHandler := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	})

	r.Use(corsHandler)

	r.GET("/books", bookController.ListBooks)
	r.GET("/book/:id", bookController.GetBookById)
	r.POST("/book", bookController.CreateBook)
	r.PUT("/book/:id", bookController.UpdateBookById)
	r.DELETE("/book/:id", bookController.DeleteBookById)

	r.GET("/health_check", healthCheckController.HealthCheck)

	_ = r.Run(":5000")
}

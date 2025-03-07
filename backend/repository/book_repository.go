package repository

import (
	"app/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	ListBooks(ctx context.Context) ([]models.Book, error)
	GetBookById(ctx context.Context, id uuid.UUID) (*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	UpdateBookById(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error)
	DeleteBookById(ctx context.Context, id uuid.UUID) error
}

type bookRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewBookRepository(db *gorm.DB, rdb *redis.Client) BookRepository {
	return &bookRepository{db: db, rdb: rdb}
}

func (repo *bookRepository) getBookFromRedis(ctx context.Context, redisKey string) (*models.Book, error) {
	jsonData, err := repo.rdb.Get(ctx, redisKey).Result()
	if err != nil {
		return nil, err
	}

	// Parse the user data
	var book models.Book
	json.Unmarshal([]byte(jsonData), &book)
	return &book, nil
}

func (repo *bookRepository) saveBookInRedis(ctx context.Context, book models.Book) {
	// Convert user to JSON
	bookData, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error marshalling book %d: %v", book.Id, err)
	}

	// Store user in Redis (using the user ID as the key)
	key := "book:" + book.Id.String()
	err = repo.rdb.Set(ctx, key, bookData, time.Hour).Err()
	if err != nil {
		log.Printf("Error storing book %d in Redis: %v", book.Id, err)
	} else {
		log.Printf("Book %d cached in Redis", book.Id)
	}
}

func (repo *bookRepository) ListBooks(ctx context.Context) ([]models.Book, error) {
	log.Println("Listing books")
	var books []models.Book
	rdb := repo.rdb
	keys, err := rdb.Keys(ctx, "book:*").Result()
	if err != nil {
		return nil, err
	}

	if len(keys) > 0 {
		for _, key := range keys {
			book, _ := repo.getBookFromRedis(ctx, key)
			books = append(books, *book)
		}
		return books, nil
	}

	if err := repo.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *bookRepository) GetBookById(ctx context.Context, id uuid.UUID) (*models.Book, error) {
	log.Println("Get book by id")
	key := "book:" + id.String()
	book, _ := repo.getBookFromRedis(ctx, key)
	if book != nil {
		return book, nil
	}

	if err := repo.db.First(book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (repo *bookRepository) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	log.Println("Create book")
	if err := repo.db.Create(&book).Error; err != nil {
		return nil, err
	}
	repo.saveBookInRedis(ctx, *book)
	return book, nil
}

func (repo *bookRepository) UpdateBookById(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error) {
	log.Println("Update book by id")
	var existingBook models.Book

	if err := repo.db.First(&existingBook, id).Error; err != nil {
		return nil, err
	}

	existingBook.Merge(book)
	if err := repo.db.Save(&existingBook).Error; err != nil {
		return nil, err
	}
	repo.saveBookInRedis(ctx, existingBook)

	return &existingBook, nil
}

func (repo *bookRepository) DeleteBookById(ctx context.Context, id uuid.UUID) error {
	log.Println("Delete book by id")
	key := "book:" + id.String()
	repo.rdb.Del(ctx, key).Result()

	if err := repo.db.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

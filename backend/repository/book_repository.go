package repository

import (
	"app/models"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewBookRepository(db *gorm.DB, rdb *redis.Client) *BookRepository {
	return &BookRepository{db: db, rdb: rdb}
}

func (repo *BookRepository) ListBooks() []models.Book {
	log.Println("Listing books")
}

func (repo *BookRepository) GetBookById(id uuid.UUID) models.Book {
	log.Println("Get book by id")
}

func (repo *BookRepository) CreateBook() models.Book {
	log.Println("Create book")
}

func (repo *BookRepository) UpdateBookById(id uuid.UUID) models.Book {
	log.Println("Update book by id")
}

func (repo *BookRepository) DeleteBookById(id uuid.UUID) {
	log.Println("Delete book by id")
}

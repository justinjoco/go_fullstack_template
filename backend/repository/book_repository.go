package repository

import (
	"app/models"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	ListBooks() ([]models.Book, error)
	GetBookById(id uuid.UUID) (*models.Book, error)
	CreateBook(book *models.Book) (*models.Book, error)
	UpdateBookById(id uuid.UUID, book *models.Book) (*models.Book, error)
	DeleteBookById(id uuid.UUID) error
}

type bookRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewBookRepository(db *gorm.DB, rdb *redis.Client) BookRepository {
	return &bookRepository{db: db, rdb: rdb}
}

func (repo *bookRepository) ListBooks() ([]models.Book, error) {
	log.Println("Listing books")
	var books []models.Book

	if err := repo.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *bookRepository) GetBookById(id uuid.UUID) (*models.Book, error) {
	log.Println("Get book by id")
	var book models.Book
	if err := repo.db.First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *bookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	log.Println("Create book")
	if err := repo.db.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (repo *bookRepository) UpdateBookById(id uuid.UUID, book *models.Book) (*models.Book, error) {
	log.Println("Update book by id")
	var existingBook models.Book

	if err := repo.db.First(&existingBook, id).Error; err != nil {
		return nil, err
	}

	existingBook.Merge(book)
	if err := repo.db.Save(&existingBook).Error; err != nil {
		return nil, err
	}

	return &existingBook, nil
}

func (repo *bookRepository) DeleteBookById(id uuid.UUID) error {
	log.Println("Delete book by id")

	if err := repo.db.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

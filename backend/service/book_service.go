package service

import (
	"app/models"
	"app/repository"
	"log"

	"github.com/google/uuid"
)

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repository: repo}
}

func (svc *BookService) ListBooks() []models.Book {
	log.Println("Listing books")
}

func (svc *BookService) GetBookById(id uuid.UUID) models.Book {
	log.Println("Get book by id")
}

func (svc *BookService) CreateBook() models.Book {
	log.Println("Create book")
}

func (svc *BookService) UpdateBookById(id uuid.UUID) models.Book {
	log.Println("Update book by id")
}

func (svc *BookService) DeleteBookById(id uuid.UUID) {
	log.Println("Delete book by id")
}

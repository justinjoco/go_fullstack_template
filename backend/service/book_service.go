package service

import (
	"app/models"
	"app/repository"
	"log"

	"github.com/google/uuid"
)

type BookService interface {
	ListBooks() []models.Book
	GetBookById(id uuid.UUID) *models.Book
	CreateBook(book *models.Book) *models.Book
	UpdateBookById(id uuid.UUID, book *models.Book) *models.Book
	DeleteBookById(id uuid.UUID)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (svc *bookService) ListBooks() []models.Book {
	log.Println("Listing books")
	books, err := svc.repo.ListBooks()
	if err != nil {
		log.Println("Error grabbing books from repo")
	}
	return books
}

func (svc *bookService) GetBookById(id uuid.UUID) *models.Book {
	log.Println("Get book by id")
	book, err := svc.repo.GetBookById(id)
	if err != nil {
		log.Println("Error grabbing a book")
	}
	return book
}

func (svc *bookService) CreateBook(book *models.Book) *models.Book {
	log.Println("Create book")
	book, err := svc.repo.CreateBook(book)
	if err != nil {
		log.Println("Error creating book in repo")
	}
	return book
}

func (svc *bookService) UpdateBookById(id uuid.UUID, book *models.Book) *models.Book {
	log.Println("Update book by id")
	book, err := svc.repo.UpdateBookById(id, book)
	if err != nil {
		log.Println("Error updating book")
	}
	return book
}

func (svc *bookService) DeleteBookById(id uuid.UUID) {
	log.Println("Delete book by id")
	err := svc.repo.DeleteBookById(id)
	if err != nil {
		log.Println("Error deleting book")
	}
}

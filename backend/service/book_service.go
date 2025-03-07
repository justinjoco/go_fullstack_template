package service

import (
	"app/models"
	"app/repository"
	"context"
	"log"

	"github.com/google/uuid"
)

type BookService interface {
	ListBooks(ctx context.Context) []models.Book
	GetBookById(ctx context.Context, id uuid.UUID) *models.Book
	CreateBook(ctx context.Context, book *models.Book) *models.Book
	UpdateBookById(ctx context.Context, id uuid.UUID, book *models.Book) *models.Book
	DeleteBookById(ctx context.Context, id uuid.UUID)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (svc *bookService) ListBooks(ctx context.Context) []models.Book {
	log.Println("Listing books")
	books, err := svc.repo.ListBooks(ctx)
	if err != nil {
		log.Println("Error grabbing books from repo")
	}
	return books
}

func (svc *bookService) GetBookById(ctx context.Context, id uuid.UUID) *models.Book {
	log.Println("Get book by id")
	book, err := svc.repo.GetBookById(ctx, id)
	if err != nil {
		log.Println("Error grabbing a book")
	}
	return book
}

func (svc *bookService) CreateBook(ctx context.Context, book *models.Book) *models.Book {
	log.Println("Create book")
	book, err := svc.repo.CreateBook(ctx, book)
	if err != nil {
		log.Println("Error creating book in repo")
	}
	return book
}

func (svc *bookService) UpdateBookById(ctx context.Context, id uuid.UUID, book *models.Book) *models.Book {
	log.Println("Update book by id")
	book, err := svc.repo.UpdateBookById(ctx, id, book)
	if err != nil {
		log.Println("Error updating book")
	}
	return book
}

func (svc *bookService) DeleteBookById(ctx context.Context, id uuid.UUID) {
	log.Println("Delete book by id")
	err := svc.repo.DeleteBookById(ctx, id)
	if err != nil {
		log.Println("Error deleting book")
	}
}

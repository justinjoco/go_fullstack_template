package service

import "app/repository"

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repository: repo}
}

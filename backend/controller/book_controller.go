package controller

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service *service.BookService
}

func NewBookController(svc *service.BookService) *BookController {
	return &BookController{service: svc}
}

func (c *BookController) ListBooks(ctx *gin.Context) {

}

func (c *BookController) GetBookById(ctx *gin.Context) {

}

func (c *BookController) CreateBook(ctx *gin.Context) {

}

func (c *BookController) UpdateBookById(ctx *gin.Context) {

}

func (c *BookController) DeleteBookById(ctx *gin.Context) {

}

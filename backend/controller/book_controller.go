package controller

import (
	"app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service *service.BookService
}

func NewBookController(svc *service.BookService) *BookController {
	return &BookController{service: svc}
}

func (c *BookController) ListBooks(ctx *gin.Context) {
	log.Println("Listing books")
	ctx.Status(http.StatusOK)

}

func (c *BookController) GetBookById(ctx *gin.Context) {
	log.Println("Get book by id")
	ctx.Status(http.StatusOK)
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	log.Println("Create book")
	ctx.Status(http.StatusCreated)

}

func (c *BookController) UpdateBookById(ctx *gin.Context) {
	log.Println("Update book by id")
	ctx.Status(http.StatusOK)

}

func (c *BookController) DeleteBookById(ctx *gin.Context) {
	log.Println("Delete book by id")
	ctx.Status(http.StatusNoContent)
}

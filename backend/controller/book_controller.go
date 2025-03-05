package controller

import (
	"app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	service *service.BookService
}

func NewBookController(svc *service.BookService) *BookController {
	return &BookController{service: svc}
}

func (c *BookController) ListBooks(ctx *gin.Context) {
	log.Println("Listing books")
	books := c.service.ListBooks()
	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) GetBookById(ctx *gin.Context) {
	log.Println("Get book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	book := c.service.GetBookById(uuidId)
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	log.Println("Create book")
	book := c.service.CreateBook()
	ctx.JSON(http.StatusCreated, book)
}

func (c *BookController) UpdateBookById(ctx *gin.Context) {
	log.Println("Update book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	book := c.service.UpdateBookById(uuidId)
	ctx.JSON(http.StatusOK, book)

}

func (c *BookController) DeleteBookById(ctx *gin.Context) {
	log.Println("Delete book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	c.service.DeleteBookById(uuidId)
	ctx.Status(http.StatusNoContent)
}

package controller

import (
	"app/models"
	"app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController interface {
	ListBooks(ctx *gin.Context)
	GetBookById(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	UpdateBookById(ctx *gin.Context)
	DeleteBookById(ctx *gin.Context)
}

type bookController struct {
	service service.BookService
}

func NewBookController(svc service.BookService) BookController {
	return &bookController{service: svc}
}

func (c *bookController) ListBooks(ctx *gin.Context) {
	log.Println("Listing books")
	books := c.service.ListBooks(ctx.Request.Context())
	ctx.JSON(http.StatusOK, books)
}

func (c *bookController) GetBookById(ctx *gin.Context) {
	log.Println("Get book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	book := c.service.GetBookById(ctx.Request.Context(), uuidId)
	ctx.JSON(http.StatusOK, book)
}

func (c *bookController) CreateBook(ctx *gin.Context) {
	log.Println("Create book")
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ret := c.service.CreateBook(ctx.Request.Context(), &book)
	ctx.JSON(http.StatusCreated, *ret)
}

func (c *bookController) UpdateBookById(ctx *gin.Context) {
	log.Println("Update book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ret := c.service.UpdateBookById(ctx.Request.Context(), uuidId, &book)
	ctx.JSON(http.StatusOK, *ret)

}

func (c *bookController) DeleteBookById(ctx *gin.Context) {
	log.Println("Delete book by id")
	id := ctx.Param("id")
	uuidId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be UUID formatted"})
	}
	c.service.DeleteBookById(ctx.Request.Context(), uuidId)
	ctx.Status(http.StatusNoContent)
}

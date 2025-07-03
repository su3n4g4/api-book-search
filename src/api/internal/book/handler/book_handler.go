package handler

import (
	"net/http"

	"api-book-search/internal/book/usecase"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}

// GET /api/books?q=golang
func (h *BookHandler) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	ctx := c.Request.Context()

	books, err := h.bookUsecase.SearchBooks(ctx, query)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, books)
}

// GET /api/books/:id
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book id is required"})
		return
	}

	book, err := h.bookUsecase.GetBookByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

package repository

import (
	"api-book-search/internal/book/entity"
)

type BookRepository interface {
	SearchBooks(query string) ([]*entity.Book, error)
	GetBookByID(volumeID string) (*entity.Book, error)
}

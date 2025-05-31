package usecase

import (
	"api-book-search/internal/book/entity"
)

type BookUsecase interface {
	SearchBooks(query string) ([]*entity.Book, error)
	GetBookByID(id string) (*entity.Book, error)
}

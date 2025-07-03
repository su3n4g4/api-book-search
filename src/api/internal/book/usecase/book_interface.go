package usecase

import (
	"api-book-search/internal/book/entity"
	"context"
)

type BookUsecase interface {
	SearchBooks(ctx context.Context, query string) ([]*entity.Book, error)
	GetBookByID(ctx context.Context, id string) (*entity.Book, error)
}

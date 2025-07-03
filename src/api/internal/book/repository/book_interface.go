package repository

import (
	"api-book-search/internal/book/entity"
	"context"
)

type BookRepository interface {
	SearchBooks(ctx context.Context, query string) ([]*entity.Book, error)
	GetBookByID(ctx context.Context, volumeID string) (*entity.Book, error)
}

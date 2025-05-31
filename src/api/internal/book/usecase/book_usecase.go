package usecase

import (
	"api-book-search/internal/book/entity"
	"api-book-search/internal/book/repository"
)

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (u *bookUsecase) SearchBooks(query string) ([]*entity.Book, error) {
	return u.bookRepo.SearchBooks(query)
}

func (u *bookUsecase) GetBookByID(id string) (*entity.Book, error) {
	return u.bookRepo.GetBookByID(id)
}

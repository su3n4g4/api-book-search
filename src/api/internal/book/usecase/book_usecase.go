package usecase

import (
	"api-book-search/internal/apperrors"
	"api-book-search/internal/book/entity"
	"api-book-search/internal/book/repository"
	"context"
	"errors"
)

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (u *bookUsecase) SearchBooks(ctx context.Context, query string) ([]*entity.Book, error) {
	if query == "" {
		return nil, apperrors.New(apperrors.ValidationError, "検索キーワードを入力してください", nil)
	}
	books, err := u.bookRepo.SearchBooks(ctx, query)
	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			switch appErr.Type {
			case apperrors.NotFoundError:
				return []*entity.Book{}, nil
			case apperrors.TimeoutError:
				// 外部APIの一時的な不具合
				return nil, apperrors.New(apperrors.TimeoutError, "外部サービスがタイムアウトしました", err)
			default:
				// それ以外はそのまま返却
				return nil, appErr
			}
		}
		// AppError以外の想定外のエラー
		return nil, apperrors.New(apperrors.InternalError, "予期しないエラーが発生しました", err)
	}

	return books, nil
}

func (u *bookUsecase) GetBookByID(ctx context.Context, id string) (*entity.Book, error) {
	return u.bookRepo.GetBookByID(ctx, id)
}

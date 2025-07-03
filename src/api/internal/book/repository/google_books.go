package repository

import (
	"api-book-search/internal/apperrors"
	"api-book-search/internal/book/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type googleBooksRepository struct {
	apiBase string
	apiKey  string
}

func NewGoogleBooksRepository(apiKey string) BookRepository {
	return &googleBooksRepository{
		apiBase: "https://www.googleapis.com/books/v1",
		apiKey:  apiKey,
	}
}

func (r *googleBooksRepository) SearchBooks(ctx context.Context, query string) ([]*entity.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	url := fmt.Sprintf("%s/volumes?q=%s", r.apiBase, url.QueryEscape(query))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, apperrors.New(apperrors.TimeoutError, "failed to build request", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.Is(err, context.DeadlineExceeded) ||
			(errors.As(err, &netErr) && netErr.Timeout()) {
			return nil, apperrors.New(apperrors.TimeoutError, "external API timeout", err)
		}
		return nil, apperrors.New(apperrors.ExternalAPIError, "failed to call external API", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, apperrors.New(apperrors.ExternalAPIError,
			fmt.Sprintf("external API error: status %d", resp.StatusCode), nil)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, apperrors.New(apperrors.NotFoundError,
			fmt.Sprintf("no results found (status: %d)", resp.StatusCode), nil)
	}

	var result struct {
		Items []struct {
			ID         string `json:"id"`
			VolumeInfo struct {
				Title     string   `json:"title"`
				Authors   []string `json:"authors"`
				Thumbnail string   `json:"imageLinks.thumbnail"`
			} `json:"volumeInfo"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, apperrors.New(apperrors.InternalError, "failed to decode API response", err)
	}

	var books []*entity.Book
	for _, item := range result.Items {
		books = append(books, &entity.Book{
			ID:        item.ID,
			Title:     item.VolumeInfo.Title,
			Authors:   item.VolumeInfo.Authors,
			Thumbnail: item.VolumeInfo.Thumbnail,
		})
	}

	return books, nil
}

func (r *googleBooksRepository) GetBookByID(ctx context.Context, id string) (*entity.Book, error) {
	url := fmt.Sprintf("%s/volumes/%s", r.apiBase, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var item struct {
		ID         string `json:"id"`
		VolumeInfo struct {
			Title      string   `json:"title"`
			Authors    []string `json:"authors"`
			ImageLinks struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}

	return &entity.Book{
		ID:        item.ID,
		Title:     item.VolumeInfo.Title,
		Authors:   item.VolumeInfo.Authors,
		Thumbnail: item.VolumeInfo.ImageLinks.Thumbnail,
	}, nil
}

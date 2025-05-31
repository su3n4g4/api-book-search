package repository

import (
	"api-book-search/internal/book/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (r *googleBooksRepository) SearchBooks(query string) ([]*entity.Book, error) {
	url := fmt.Sprintf("%s/volumes?q=%s", r.apiBase, url.QueryEscape(query))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
		return nil, err
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

func (r *googleBooksRepository) GetBookByID(id string) (*entity.Book, error) {
	url := fmt.Sprintf("%s/volumes/%s", r.apiBase, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var item struct {
		ID         string `json:"id"`
		VolumeInfo struct {
			Title     string   `json:"title"`
			Authors   []string `json:"authors"`
			Thumbnail string   `json:"imageLinks.thumbnail"`
		} `json:"volumeInfo"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}

	return &entity.Book{
		ID:        item.ID,
		Title:     item.VolumeInfo.Title,
		Authors:   item.VolumeInfo.Authors,
		Thumbnail: item.VolumeInfo.Thumbnail,
	}, nil
}

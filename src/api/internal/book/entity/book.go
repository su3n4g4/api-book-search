package entity

type Book struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Authors   []string `json:"authors"`
	Thumbnail string   `json:"imageLinks.thumbnail"`
}

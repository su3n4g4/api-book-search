package main

import (
	"api-book-search/infra"
	"api-book-search/models"
)

func main() {
    infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}
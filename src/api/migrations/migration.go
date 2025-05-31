package main

import (
	"api-book-search/internal/memo/entity"
	"api-book-search/pkg/infrastructure"
	"fmt"
	"log"
	"os"
)

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
	}

	db, err := infrastructure.NewDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&entity.Memo{}); err != nil {
		panic("failed to migrate database")
	}
}

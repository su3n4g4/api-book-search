package main

import (
	"api-book-search/internal/book/handler"
	"api-book-search/internal/book/repository"
	"api-book-search/internal/book/usecase"
	"api-book-search/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	// --- DB接続 ---
	// db, err := infrastructure.NewDB(dsn)
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// db := infrastructure.SetupDB()

	// 環境変数の読み込み
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルト
	}

	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	// DI
	bookRepo := repository.NewGoogleBooksRepository(apiKey)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	// ルーティング設定
	api := router.Group("/api")
	{
		api.GET("/books", bookHandler.SearchBooks)
		api.GET("/books/:id", bookHandler.GetBookByID)
	}

	router.Run(":8080")
}

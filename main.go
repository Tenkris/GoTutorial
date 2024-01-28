package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func loggingMiddleware(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate processing time
	duration := time.Since(start)

	// Log the information
	fmt.Printf("Request URL: %s - Method: %s - Duration: %s\n", c.OriginalURL(), c.Method(), duration)

	return err
}

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Get secret key from environment
	secretKey := os.Getenv("SECRET_KEY")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/login", login(secretKey))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secretKey),
	}))

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang"})
	books = append(books, Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine"})
	books = append(books, Book{ID: 3, Title: "Golang routers", Author: "Mr. Router"})
	app.Use(loggingMiddleware)
	app.Get("/api/v1/books", getBooks)
	app.Get("/api/v1/book/:id", getBook)

	app.Use(IsAdmin)
	app.Post("/api/v1/book", createBook)
	app.Put("/api/v1/book/:id", updateBook)
	app.Delete("/api/v1/book/:id", deleteBook)
	app.Post("/api/v1/upload", uploadImage)
	app.Get("/api/v1/env", getConfig)
	app.Listen(":3000")
}

func getConfig(c *fiber.Ctx) error {
	return c.JSON(os.Getenv("PORT"))
}

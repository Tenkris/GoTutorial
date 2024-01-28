package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
func getBook(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
func createBook(c *fiber.Ctx) error {
	book := new(Book)
	err := c.BodyParser(&book)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book.ID = len(books) + 1
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book := new(Book)
	err = c.BodyParser(&book)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book.ID = id
	for i, b := range books {
		if b.ID == id {
			b.Author = book.Author
			b.Title = book.Title
			books[i] = b
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
func deleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

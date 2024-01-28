package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Dummy user for example
var user = struct {
	Email    string
	Password string
}{
	Email:    "user@example.com",
	Password: "password123",
}

func login(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var request LoginRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		fmt.Println(request)

		// Check credentials - In real world, you should check against a database
		if request.Email != user.Email || request.Password != user.Password {
			return fiber.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "John Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}
}
func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	if user == nil || user.Claims.(jwt.MapClaims)["admin"] != true {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}

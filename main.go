package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Firstname string
	Lastname  string
	Id        string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "John",
		Lastname:  "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.New().String()

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	userGroup := app.Group("/user")

	userGroup.Get("/user", handleUser)
	userGroup.Post("/user", handleCreateUser)

	app.Listen(":3000")
}

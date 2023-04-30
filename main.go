package main

import (
	"github.com/andrevks/simple-crm-go-fiber/database"
	"github.com/andrevks/simple-crm-go-fiber/user"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	// Create a /api/v1 endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
				"message": "Welcome to the Simple CRM API!",
				"endpoints": fiber.Map{
						"Get all users": "/api/v1/users",
						"Get user by ID": "/api/v1/users/:id",
						"Create new user": "/api/v1/users",
						"Delete user": "/api/v1/users/:id",
				},
		})	
	})
	app.Get("/api/v1/users", user.GetUsers)
	app.Get("/api/v1/users/:id", user.GetUser)
	app.Post("/api/v1/users", user.NewUser)
	app.Delete("/api/v1/users/:id",user.DeleteUser)
}

func main() {
	app := fiber.New()
	database.Connect()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.Disconnect()
}
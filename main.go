package main

import (
	"github.com/andrevks/simple-crm-go-fiber/database"
	"github.com/andrevks/simple-crm-go-fiber/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	database.Connect()
	routes.SetupRoutes(app)
	app.Listen(":3000")
	defer database.Disconnect()
}
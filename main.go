package main

import (
	"fmt"
	"github.com/andrevks/simple-crm-go-fiber/database"
	"github.com/andrevks/simple-crm-go-fiber/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	// Create a /api/v1 endpoint
	app.Get("/api/v1/users", user.GetUsers)
	app.Get("/api/v1/users/:id", user.GetUser)
	app.Post("/api/v1/users", user.NewUser)
	app.Delete("/api/v1/users/:id",user.DeleteUser)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "simple-crm.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}
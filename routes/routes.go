package routes

import(
	"github.com/andrevks/simple-crm-go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
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

	api.Get("/users", user.GetUsers)
	api.Get("/users/:id", user.GetUser)
	api.Post("/users", user.NewUser)
	api.Delete("/users/:id", user.DeleteUser)
}
package user

import(
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/andrevks/simple-crm-go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/andrevks/simple-crm-go-fiber/models"
)
 

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.FindAll(&users, "id DESC") 
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx)error {
	id := c.Params("id")
	var user models.User 
	database.Find(&user, id)
	return c.JSON(user)
}

func NewUser(c *fiber.Ctx) error{
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	database.Create(user)
	return c.JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.Find(&user, id)
	if user.Name == "" {
		return c.Status(500).Send([]byte("No lead found with ID"))
	}
	database.Delete(&user,id)
	return c.SendStatus(204)
}
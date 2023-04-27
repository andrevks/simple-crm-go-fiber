package user

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/andrevks/simple-crm-go-fiber/database"
	"github.com/gofiber/fiber/v2"
)
 
type User struct {
	gorm.Model
	Name				string	`json:"name"`
	Company			string  `json:"company"`
	Email				string  `json:"email"`
	Phone				int 		`json:"phone"`
}


func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx)error {
	id := c.Params("id")
	db := database.DBConn
	var user User 
	db.Find(&user, id)
	return c.JSON(user)
}

func NewUser(c *fiber.Ctx) error{
	db := database.DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
		 
	}

	db.Create(&user)
	return c.JSON(user)
}


func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.First(&user, id)
	if user.Name == "" {
		return c.Status(500).Send([]byte("No lead found with ID"))
	}

	db.Delete(&user)
	return c.SendStatus(204)
}
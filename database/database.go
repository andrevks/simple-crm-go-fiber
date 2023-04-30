package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/andrevks/simple-crm-go-fiber/models"
)

var(
	DBConn *gorm.DB
	dbPath = "./simple-crm.db"
)

// Connect with database
func Connect() {
	var err error
	DBConn, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection opened to database")
	DBConn.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}

// Disconnect from database
func Disconnect() {
	DBConn.Close()
	fmt.Println("Connection closed to database")
}

// basic CRUD operations with GORM and SQLite for any	Model

func FindAll(modelSlice interface {}, order string )   {
	DBConn.Order(order).Find(modelSlice)
}

func Create(model models.Model)  {
	DBConn.Create(model)
}

func Find(model models.Model, id string)  {
	DBConn.First(model, id)
}
	
func Delete(model models.Model, id string) {
	DBConn.First(model, id)
	DBConn.Delete(model)
}
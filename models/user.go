package models

import(
	"github.com/jinzhu/gorm"
)
 
type User struct {
	gorm.Model
	Name				string	`json:"name"`
	Company			string  `json:"company"`
	Email				string  `json:"email"`
	Phone				int 		`json:"phone"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) SetID(id uint) {
	u.ID = id
}
package usermodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	Phonenumber string
	Storename string
	Role string `json:"role"`
}

func SelectAll() *gorm.DB {
	items := []User{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item User
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *User) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdatesCustomer(id string, newCustomer *User) *gorm.DB {
	var item User
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCustomer)
}

func UpdatesSeller(id string, newSeller *User) *gorm.DB {
	var item User
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newSeller)
}

func Deletes(id string) *gorm.DB {
	var item User
	return config.DB.Delete(&item, "id = ?", id)
}

func FindEmail(input *User) []User {
	items := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items
}
//testing
func FindRole(input *User) []User  {
	items := []User{}
	config.DB.Raw("SELECT role FROM users WHERE email = ?", input.Role).Scan(&items)
	return items
}

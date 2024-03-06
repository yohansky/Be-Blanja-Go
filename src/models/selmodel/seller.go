package selmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Seller struct {
	gorm.Model
	SName        string
	Email        string
	Phone_Number string
	Store_Name   string
	Password     string
}

func SelectAll() *gorm.DB {
	items := []Seller{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Seller
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Seller) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newSeller *Seller) *gorm.DB {
	var item Seller
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newSeller)
}

func Deletes(id string) *gorm.DB {
	var item Seller
	return config.DB.Delete(&item, "id = ?", id)
}
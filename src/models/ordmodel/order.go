package ordmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Productid  uint
	Name       string
	Price      float64
	Sum        float64
	Delivery   float64
	Total      float64
	Costumerid uint
	CName      string
}

func SelectAll() *gorm.DB {
	items := []Order{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Order
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Order) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newOrder *Order) *gorm.DB {
	var item Order
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newOrder)
}

func Deletes(id string) *gorm.DB {
	var item Order
	return config.DB.Delete(&item, "id = ?", id)
}
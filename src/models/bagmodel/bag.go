package bagmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Bag struct {
	gorm.Model
	Productid uint
	Total_Price float64
}

func SelectAll() *gorm.DB {
	items := []Bag{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB  {
	var item Bag
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Bag) *gorm.DB  {
	return config.DB.Create(&item)
}

func Updates(id string, newBag *Bag) *gorm.DB {
	var item Bag
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newBag)
}

func Deletes(id string) *gorm.DB {
	var item Bag
	return config.DB.Delete(&item, "id = ?", id)
}
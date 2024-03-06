package addmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Alias string
	RName string
	RPhone string
	Street string
	Postal string
	City string
}

func SelectAll() *gorm.DB {
	items := []Address{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB  {
	var item Address
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Address) *gorm.DB  {
	return config.DB.Create(&item)
}

func Updates(id string, newAddress *Address) *gorm.DB {
	var item Address
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newAddress)
}

func Deletes(id string) *gorm.DB {
	var item Address
	return config.DB.Delete(&item, "id = ?", id)
}
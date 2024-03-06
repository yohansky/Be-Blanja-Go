package cosmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Costumer struct {
	gorm.Model
	CName    string
	Email    string
	Password string
}

func SelectAll() *gorm.DB {
	items := []Costumer{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB  {
	var item Costumer
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Costumer) *gorm.DB  {
	return config.DB.Create(&item)
}

func Updates(id string, newCostumer *Costumer) *gorm.DB {
	var item Costumer
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCostumer)
}

func Deletes(id string) *gorm.DB {
	var item Costumer
	return config.DB.Delete(&item, "id = ?", id)
}
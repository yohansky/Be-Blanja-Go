package paymodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Payment struct {
	gorm.Model
	Method string
	Checkid uint
}

func SelectAll() *gorm.DB {
	items := []Payment{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB  {
	var item Payment
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Payment) *gorm.DB  {
	return config.DB.Create(&item)
}

func Updates(id string, newPayment *Payment) *gorm.DB {
	var item Payment
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newPayment)
}

func Deletes(id string) *gorm.DB {
	var item Payment
	return config.DB.Delete(&item, "id = ?", id)
}
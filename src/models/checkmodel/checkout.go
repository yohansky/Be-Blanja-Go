package checkmodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Checkout struct {
	gorm.Model
	Userid uint
	Bagid uint
	Addressid uint
	Total float64
}

func SelectAll() *gorm.DB {
	items := []Checkout{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB  {
	var item Checkout
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Checkout) *gorm.DB  {
	return config.DB.Create(&item)
}

func Updates(id string, newCheckout *Checkout) *gorm.DB {
	var item Checkout
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCheckout)
}

func Deletes(id string) *gorm.DB {
	var item Checkout
	return config.DB.Delete(&item, "id = ?", id)
}
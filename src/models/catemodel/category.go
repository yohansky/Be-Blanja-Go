package catemodel

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	Imgurl string
	// Products []Product
}

// type Product struct {
// 	gorm.Model
// 	Name        string
// 	Rating      uint
// 	Price       float64
// 	Color       string
// 	Size        string
// 	Amount      uint16
// 	Condition   string
// 	Description string
// 	Sellerid    uint
// 	Imgurl      string
// 	CategoryID  uint
// }

func SelectAllCategory() *gorm.DB {
	items := []Category{}
	return config.DB.Find(&items)
}

func SelectCategory(id string) *gorm.DB {
	var item Category
	return config.DB.First(&item, "id = ?", id)
}

func PostCategory(item *Category) *gorm.DB {
	return config.DB.Create(&item)
}

func UpdatesCategory(id string, newCategory *Category) *gorm.DB {
	var item Category
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newCategory)
}

func DeletesCategory(id string) *gorm.DB {
	var item Category
	return config.DB.Delete(&item, "id = ?", id)
}

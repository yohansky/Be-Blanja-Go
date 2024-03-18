package prodmodels

import (
	"Backend-Golang/src/config"

	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Rating      uint
	Price       float64
	Color       string
	Size        string
	Amount      uint16
	Condition   string
	Description string
	Sellerid    uint
	Imgurl      string
	CategoryId  uint
	Category    Category `gorm:"foreignKey:CategoryId"`
}

type Category struct {
	gorm.Model
	Name   string
	Imgurl string
}

func SelectAllProduct() []*Product {
	items := []*Product{}
	config.DB.Find(&items)
	return items
	// return config.DB.Preload("Category").Find(&items)
}

func SelectProduct(id string) *Product {
	var item Product
	if err := config.DB.Preload("Category").First(&item, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return nil
	}
	// config.DB.Preload("Category").First(&item, "id = ?", id)
	return &item
}

func PostProduct(item *Product) *Product {
	config.DB.Create(&item)
	return item
}

func UpdatesProduct(id string, newProduct *Product) *Product {
	var item Product
	config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
	return &item
}

func DeletesProduct(id string) {
	var item Product
	config.DB.Delete(&item, "id = ?", id)
}

func FindCondProduct(sort string, limit int, offset int) []*Product {
	items := []*Product{}
	config.DB.Order(sort).Limit(limit).Offset(offset).Preload("Category").Find(&items)
	return items
}

func CountDataProduct() int64 {
	var result int64
	config.DB.Table("products").Count(&result)
	return result
}

func FindDataProduct(name string) []*Product {
	items := []*Product{}
	name = "%" + name + "%"
	config.DB.Where("name LIKE ?", name).Find(&items)
	return items
}

// func SelectAllProduct() *gorm.DB {
// 	items := []Product{}
// 	return config.DB.Find(&items)
// 	// return config.DB.Preload("Category").Find(&items)
// }

// func SelectProduct(id string) *gorm.DB {
// 	var item Product
// 	return config.DB.Preload("Category").First(&item, "id = ?", id)
// }

// func PostProduct(item *Product) *gorm.DB {
// 	return config.DB.Create(&item)
// }

// func UpdatesProduct(id string, newProduct *Product) *gorm.DB {
// 	var item Product
// 	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
// }

// func DeletesProduct(id string) *gorm.DB {
// 	var item Product
// 	return config.DB.Delete(&item, "id = ?", id)
// }

// func FindCondProduct(sort string, limit int, offset int) *gorm.DB {
// 	items := []Product{}
// 	return config.DB.Order(sort).Limit(limit).Offset(offset).Preload("Category").Find(&items)
// }

// func CountDataProduct() int {
// 	var result int
// 	config.DB.Table("products").Count(&result)
// 	return result
// }

// func FindDataProduct(name string) *gorm.DB {
// 	items := []Product{}
// 	name = "%" + name + "%"
// 	return config.DB.Where("name LIKE ?", name).Find(&items)
// }

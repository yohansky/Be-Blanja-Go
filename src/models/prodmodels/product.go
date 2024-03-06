package prodmodels

import (
	"Backend-Golang/src/config"

	"github.com/jinzhu/gorm"
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
}

// type Costumer_User struct {
// 	gorm.Model
// 	Costumerid uint
// 	Email string
// 	Password string
// }

// type Seller_User struct {
// 	gorm.Model
// 	Sellerid uint
// 	Email string
// 	Password string
// }

func SelectAll() *gorm.DB {
	items := []Product{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Product) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newProduct *Product) *gorm.DB {
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes(id string) *gorm.DB {
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}

func FindCond(sort string, limit int, offset int) *gorm.DB {
	items := []Product{}
	return config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
}

func CountData() int {
	var result int
	config.DB.Table("products").Count(&result)
	return result
}

func FindData(name string) *gorm.DB {
	items := []Product{}
	name = "%" + name + "%"
	return config.DB.Where("name LIKE ?", name).Find(&items)
}

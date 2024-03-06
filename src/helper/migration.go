package helper

import (
	"Backend-Golang/src/config"
	"Backend-Golang/src/models/addmodel"
	"Backend-Golang/src/models/bagmodel"
	"Backend-Golang/src/models/checkmodel"
	"Backend-Golang/src/models/cosmodel"
	"Backend-Golang/src/models/paymodel"
	"Backend-Golang/src/models/prodmodels"
	"Backend-Golang/src/models/selmodel"
	"Backend-Golang/src/models/usermodel"
)

func Migrate()  {
	config.DB.AutoMigrate(&cosmodel.Costumer{})
	// config.DB.AutoMigrate(&ordmodel.Order{})
	config.DB.AutoMigrate(&selmodel.Seller{})
	config.DB.AutoMigrate(&prodmodels.Product{})
	config.DB.AutoMigrate(&usermodel.User{})
	config.DB.AutoMigrate(&addmodel.Address{}) 
	config.DB.AutoMigrate(&bagmodel.Bag{})
	config.DB.AutoMigrate(&checkmodel.Checkout{})
	config.DB.AutoMigrate(&paymodel.Payment{})
}
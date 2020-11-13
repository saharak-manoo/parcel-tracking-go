package db

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/saharak/parcel-tracking-go/app/models"
)

type Seeder struct {}

var couriers = []models.Courier{
	{
		Name: "เคอรี่ เอ็กซ์เพรส",
		Color: "#FB8523",
		Active: true,
		ImageUrl: "https://etrackings.com/courier/kerry_express.jpg",
	},
	{
		Name: "แฟลช เอ็กซ์เพรส",
		Color: "#FFD21A",
		Active: true,
		ImageUrl: "https://etrackings.com/courier/flash_express.png",
	},
	{
		Name: "เจแอนด์ที เอ็กซ์เพรส",
		Color: "#FE1B3A",
		Active: true,
		ImageUrl: "https://etrackings.com/courier/jt_express.jpg",
	},
}

func (s *Seeder) Load(db *gorm.DB) {
	for i, _ := range couriers {
		err := db.Debug().Model(&models.User{}).Create(&couriers[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
	}
}
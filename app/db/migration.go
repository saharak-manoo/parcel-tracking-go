package db

import (
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/saharak/parcel-tracking-go/app/models"
)

type Migration struct {}

func (m *Migration) Load(db *gorm.DB) {
	isDropTableIfExists, err := strconv.ParseBool(os.Getenv("DB_DROP_TABLE_IF_EXISTS"))
	if err != nil {
		log.Fatalf("Cannot drop table")
	} else {
		if isDropTableIfExists == true {
			err := db.Debug().DropTableIfExists(
				&models.Courier{},
				&models.TrackHistory{},
				&models.User{},
			).Error

			if err != nil {
				log.Fatalf("Cannot drop table: %v", err)
			}
		}
	}

	err = db.Debug().AutoMigrate(
		&models.Courier{},
		&models.TrackHistory{},
		&models.User{},
	).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}
}
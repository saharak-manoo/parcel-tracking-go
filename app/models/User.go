package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UUID   		      string         `gorm:"size:100;not null;unique" json:"uuid"`
	Name   		      string         `gorm:"size:255;not null;unique" json:"name"`
	Email           string         `gorm:"size:100;not null;unique" json:"email"`
	ImageUrl        string         `gorm:"size:100;not null;" json:"imageUrl"`
	TrackHistories  []TrackHistory `json:"trackHistories"`
}
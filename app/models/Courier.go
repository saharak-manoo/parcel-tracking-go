package models

import (
	"github.com/jinzhu/gorm"
)

type Courier struct {
	gorm.Model
	Name   		      string         `gorm:"size:255;not null;unique" json:"name"`
	Color           string         `gorm:"size:20;not null;unique" json:"color"`
	Active          bool           `gorm:"default:true;not null" json:"active"`
	ImageUrl        string         `gorm:"size:100;not null;" json:"imageUrl"`
	TrackHistories  []TrackHistory `json:"trackHistories"`
}
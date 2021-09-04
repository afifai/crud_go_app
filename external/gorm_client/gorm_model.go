package gorm_client

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Slug        string `gorm:"size:255;not null;unique"`
	Description string
	Duration    int    `gorm:"size:5"`
	Image       string `gorm:"size:255"`
}

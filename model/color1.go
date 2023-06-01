package model

import "gorm.io/gorm"

type Color1 struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(60);not null"`
	Description string `gorm:"type:text"`
}

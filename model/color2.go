package model

import "gorm.io/gorm"

type Color2 struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(60);not null"`
	Description string `gorm:"type:text"`
}

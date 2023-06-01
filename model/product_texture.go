package model

import "gorm.io/gorm"

type ProductTexture struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;type:varchar(20);not null"`
	Name        string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Description string `gorm:"type:text"`
}

package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm:"uniqueIndex;type:varchar(8);not null"`
	Name  string `gorm:"type:varchar(60)"`
	Image string `gorm:"type:varchar(255);not null"`
}

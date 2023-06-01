package model

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex;type:varchar(60);not null"`
	Image string `gorm:"type:varchar(255);not null"`

	ProductID uint
	Product   Product
}

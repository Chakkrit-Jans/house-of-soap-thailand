package model

import "gorm.io/gorm"

type CustomerType struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(15);not null"`
	Description string `gorm:"type:varchar(50)"`
}

package model

import "gorm.io/gorm"

type SmellGroup struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Description string `gorm:"type:varchar(120)"`
}

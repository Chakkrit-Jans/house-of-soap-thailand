package model

import (
	"gorm.io/gorm"
)

type Title struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(12);not null"`
	Description string `gorm:"type:varchar(25)"`
}

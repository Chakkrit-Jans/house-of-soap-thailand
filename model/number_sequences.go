package model

import "gorm.io/gorm"

type NumberSequences struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(25);not null"`
	Description string `gorm:"type:varchar(255)"`
	Format      string `gorm:"type:varchar(35)"`
}

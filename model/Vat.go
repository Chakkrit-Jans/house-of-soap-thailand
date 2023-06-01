package model

import "gorm.io/gorm"

type Vat struct {
	gorm.Model
	Name       string  `gorm:"uniqueIndex;type:varchar(50);not null"`
	Percentage float32 `gorm:"type:decimal(18,5)"`
}

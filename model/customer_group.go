package model

import "gorm.io/gorm"

type CustomerGroup struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(15);not null"`
	Description string `gorm:"type:varchar(50)"`

	VatID uint `gorm:"not null"`
	Vat   Vat
}

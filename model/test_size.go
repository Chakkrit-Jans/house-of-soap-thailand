package model

import "gorm.io/gorm"

type TestSize struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Qty  int    `gorm:"not null"`

	UnitsID uint
	Units   Units
}

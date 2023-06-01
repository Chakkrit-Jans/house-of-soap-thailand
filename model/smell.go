package model

import "gorm.io/gorm"

type Smell struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;type:varchar(60);not null"`
	Description string `gorm:"type:varchar(120);not null"`

	SmellExtractsID uint
	SmellExtracts   SmellExtracts
	SmellTypeID     uint
	SmellType       SmellType
	SmellGroupID    uint
	SmellGroup      SmellGroup
}

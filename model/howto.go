package model

import "gorm.io/gorm"

type Howto struct {
	gorm.Model
	Code            string `gorm:"uniqueIndex;type:varchar(12);not null"`
	Description     string `gorm:"type:text"`
	DescriptionThai string `gorm:"type:text"`
}

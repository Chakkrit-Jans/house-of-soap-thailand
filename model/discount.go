package model

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Name                string  `gorm:"uniqueIndex;type:varchar(50);not null"`
	PercentageDiscount  float32 `gorm:"type:decimal(18,5)"`
	LimitAmountDiscount float32 `gorm:"type:decimal(18,5)"`
}

package model

import (
	"hst-api/enum"

	"gorm.io/gorm"
)

// ========== Model Cost ==========
type Cost struct {
	gorm.Model
	Code      string         `gorm:"uniqueIndex;type:varchar(60);not null"`
	Name      string         `gorm:"uniqueIndex;type:varchar(60);not null"`
	CostType  enum.CostType  `gorm:"type:ENUM('Cost per Item', 'Cost per Qty')"`
	CostGroup enum.CostGroup `gorm:"type:ENUM('Production Cost', 'Packaging Cost', 'Label Cost', 'Extra freight Cost', 'Addition Cost')"`
	Amount    float32        `gorm:"type:decimal(18,5)"`
}

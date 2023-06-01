package model

import "gorm.io/gorm"

type AppConfig struct {
	gorm.Model
	ProductionCostPerItem float32 `gorm:"not null;type:decimal(18,5)"`
	ProductionCostPerUnit float32 `gorm:"not null;type:decimal(18,5)"`
	PercentageProfit      float32 `gorm:"not null;type:decimal(18,5)"`
}

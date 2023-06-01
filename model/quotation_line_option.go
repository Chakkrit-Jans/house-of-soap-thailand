package model

import (
	"hst-api/enum"

	"gorm.io/gorm"
)

type QuotationLineOption struct {
	gorm.Model
	OptionCode string `gorm:"type:varchar(125);not null"`
	OptionName string `gorm:"type:varchar(125);not null"`
	CostType   enum.CostType
	Qty        int     `gorm:"not null"`
	Price      float32 `gorm:"type:decimal(18,5)"`
	Amount     float32 `gorm:"type:decimal(18,5)"`

	CostPerUnit float32 `gorm:"type:decimal(18,5)"`
	CostAmount  float32 `gorm:"type:decimal(18,5)"`

	Profit       float32 `gorm:"type:decimal(18,5)"`
	ProfitAmount float32 `gorm:"type:decimal(18,5)"`

	QuotationLineID uint
	CostID          uint
	Cost            Cost
}

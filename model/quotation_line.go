package model

import "gorm.io/gorm"

type QuotationLine struct {
	gorm.Model
	LineNum         int    `gorm:"primaryKey:LineNum;not null"`
	ItemCode        string `gorm:"type:varchar(125);not null"`
	ItemName        string `gorm:"type:varchar(125);not null"`
	ProductName     string `gorm:"type:varchar(60)"`
	ProductTypeName string `gorm:"type:varchar(60);not null"`
	TestSizeName    string `gorm:"type:varchar(50);not null"`
	Qty             int    `gorm:"not null"`
	QtyTestSize     int
	Price           float32 `gorm:"type:decimal(18,5)"`
	Amount          float32 `gorm:"type:decimal(18,5)"`

	//---------- Cost for calculation ----------
	ItemCost              float32 `gorm:"type:decimal(18,5)"`
	ProductionCostPerItem float32 `gorm:"type:decimal(18,5)"`
	ProductionCostPerUnit float32 `gorm:"type:decimal(18,5)"`
	ProductionCostTotal   float32 `gorm:"type:decimal(18,5)"`
	AdditionAmount        float32 `gorm:"type:decimal(18,5)"`

	//---------- Profit for calculation ----------
	Profit       float32 `gorm:"type:decimal(18,5)"`
	ProfitAmount float32 `gorm:"type:decimal(18,5)"`

	//Status      int8
	QuotationID uint

	QuotationLineOption []QuotationLineOption `gorm:"foreignKey:QuotationLineID"`

	ProductID        uint
	Product          Product
	ProductTypeID    uint
	ProductType      ProductType
	ProductFormulaID uint
	ProductFormula   ProductFormula
	TestSizeID       uint
	TestSize         TestSize
}

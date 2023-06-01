package model

import (
	"hst-api/enum"
	"time"

	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	Number   string               `gorm:"uniqueIndex;type:varchar(12);not null"`
	Revision uint8                `gorm:"not null"`
	Status   enum.QuotationStatus `gorm:"type:ENUM('Draft', 'Completed', 'Approved', 'Canceled', 'Rejected')"`

	CustomerCode string    `gorm:"type:varchar(12);not null"`
	CustomerName string    `gorm:"type:varchar(160);not null"`
	DocumentDate time.Time `gorm:"type:date"`
	SalesName    string    `gorm:"type:varchar(160);not null"`

	Total              float32 `gorm:"type:decimal(18,5)"`
	DiscountPercentage float32 `gorm:"type:decimal(18,5)"`
	DiscountAmount     float32 `gorm:"type:decimal(18,5)"`
	VatPercentage      float32 `gorm:"type:decimal(18,5)"`
	VatAmount          float32 `gorm:"type:decimal(18,5)"`
	GrandTotal         float32 `gorm:"type:decimal(18,5)"`

	QuotationLine []QuotationLine `gorm:"foreignKey:QuotationID"`

	CustomersID uint
	Customers   Customers
	SalesID     uint
	Sales       Sales
	VatID       uint
	Vat         Vat
}

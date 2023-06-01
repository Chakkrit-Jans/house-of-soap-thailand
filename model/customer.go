package model

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Code      string `gorm:"uniqueIndex;type:varchar(12);not null"`
	FirstName string `gorm:"type:varchar(80);not null"`
	LastName  string `gorm:"type:varchar(80)"`
	FullName  string `gorm:"type:varchar(160);not null"`
	NickName  string `gorm:"type:varchar(25)"`
	Email     string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Mobile    string `gorm:"type:varchar(25);not null"`
	Image     string `gorm:"type:varchar(255);not null"`
	Status    uint8

	TitleID         uint
	Title           Title
	CustomerGroupID uint
	CustomerGroup   CustomerGroup
	CustomerTypeID  uint
	CustomerType    CustomerType
	VatID           uint
	Vat             Vat
}

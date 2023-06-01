package model

import "gorm.io/gorm"

type ProductFormula struct {
	gorm.Model
	Code            string  `gorm:"uniqueIndex;type:varchar(125);not null"`
	Name            string  `gorm:"uniqueIndex;type:varchar(125);not null"`
	Image           string  `gorm:"type:varchar(255);not null"`
	Properties      string  `gorm:"type:text"`
	ActiveIngedient string  `gorm:"type:text"`
	UnitCost        float32 `gorm:"type:decimal(18,5)"`
	SalePrice       float32 `gorm:"type:decimal(18,5)"`

	ProductID        uint
	Product          Product
	ProductTypeID    uint
	ProductType      ProductType
	SmellID          uint
	Smell            Smell
	ProductTextureID uint
	ProductTexture   ProductTexture
	Color1ID         uint
	Color1           Color1
	Color2ID         uint
	Color2           Color2
	Color3ID         uint
	Color3           Color3
	TestSizeID       uint
	TestSize         TestSize
	ClaimID          uint
	Claim            Claim
	HowtoID          uint
	Howto            Howto
}

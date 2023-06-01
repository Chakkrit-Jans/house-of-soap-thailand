package model

import "gorm.io/gorm"

type UserApprovedDiscount struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(60);not null"`

	SystemUsersID uint
	SystemUsers   SystemUsers
	DiscountID    uint
	Discount      Discount
}

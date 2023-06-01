package model

import "gorm.io/gorm"

type SystemUsers struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;type:varchar(50);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(80);not null"`
	LastName  string `gorm:"type:varchar(80)"`
	FullName  string `gorm:"type:varchar(160);not null"`
	NickName  string `gorm:"type:varchar(25)"`
	Mobile    string `gorm:"type:varchar(25);not null"`
	Image     string `gorm:"type:varchar(255);not null"`

	Status uint8 `gorm:"not null"`
	Verify bool  `gorm:"not null"`

	TitleID uint
	Title   Title
	SalesID uint
	Sales   Sales
}

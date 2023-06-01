package model

import (
	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	Code      string `gorm:"uniqueIndex;type:varchar(12);not null"`
	FirstName string `gorm:"type:varchar(80);not null"`
	LastName  string `gorm:"type:varchar(80)"`
	FullName  string `gorm:"type:varchar(160);not null"`
	NickName  string `gorm:"type:varchar(25)"`
	Email     string `gorm:"type:varchar(50);not null"`
	Mobile    string `gorm:"type:varchar(25);not null"`
	Image     string `gorm:"type:varchar(255);not null"`
	Status    uint8

	TitleID     uint
	Title       Title
	SalesTeamID uint
	SalesTeam   SalesTeam
}

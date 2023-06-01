package model

import (
	"time"

	"gorm.io/gorm"
)

type UsersLogs struct {
	gorm.Model
	Token            string    `gorm:"uniqueIndex;type:varchar(50);not null"`
	Name             string    `gorm:"type:varchar(255);not null"`
	SignInTime       time.Time `gorm:"datetime;not null"`
	SignOutTime      time.Time `gorm:"type:datetime"`
	IsActive         bool      `gorm:"not null"`
	DeviceID         string    `gorm:"type:varchar(120);not null"`
	DeviceName       string    `gorm:"type:varchar(120)"`
	DeviceModel      string    `gorm:"type:varchar(120)"`
	IsTablet         bool
	SystemName       string `gorm:"type:varchar(120)"`
	SystemVersion    string `gorm:"type:varchar(120)"`
	IsPhysicalDevice bool
	IPv4             string `gorm:"type:varchar(20)"`
	IPv6             string `gorm:"type:varchar(120)"`
	Event            string

	SystemUsersID uint
	SystemUsers   SystemUsers
}

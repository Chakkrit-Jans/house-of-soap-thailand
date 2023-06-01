package model

import (
	"gorm.io/gorm"
)

type UsersGroupDetail struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(50);not null"`

	SystemUsersID uint
	SystemUsers   SystemUsers
	UsersGroupID  uint
	UsersGroup    UsersGroup
}

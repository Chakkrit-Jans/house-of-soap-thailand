package model

import (
	"gorm.io/gorm"
)

type UsersGroupAuthorized struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(50);not null"`

	UsersGroupID       uint
	UsersGroup         UsersGroup
	SystemAuthorizedID uint
	SystemAuthorized   SystemAuthorized
	UserPermissionsID  uint
	UserPermissions    UserPermissions
}

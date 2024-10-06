package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Role struct {
	GormModel
	RoleName  string `gorm:"not null;uniqueIndex" json:"role_name" form:"role_name" valid:"required~Role name is required"`
	GuardName string `gorm:"not null" json:"guard_name" form:"guard_name" valid:"required~Guard name is required"`
	Users     []User `gorm:"many2many:user_roles;" json:"users"`
}

func (u *Role) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"time"

	"github.com/risdatamamal/api-javaprojects/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName      string     `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~User name is required"`
	Email         string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	Password      string     `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)" `
	IsActive      bool       `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	EmailVerified *time.Time `json:"created_at,omitempty"`
	Roles         []Role     `gorm:"many2many:user_roles;" json:"roles"`
	// Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	// Reviews     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	hashedPass := helpers.HashPass(u.Password)

	u.Password = hashedPass
	return nil
}

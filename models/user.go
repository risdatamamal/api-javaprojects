package models

import (
	"time"

	"github.com/risdatamamal/api-javaprojects/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName        string     `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~User name is required"`
	Email           string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	Password        string     `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)" `
	IsActive        bool       `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	PhotoPath       *string    `json:"photo_path,omitempty" form:"photo_path"`
	RoleID          int        `gorm:"not null" json:"role_id"`
	Role            *Role
	// Reviews         []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
}

type GetAllUsersResponse struct {
	GormModel
	UserName        string     `json:"user_name"`
	Email           string     `json:"email"`
	IsActive        bool       `json:"is_active"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	PhotoPath       *string    `json:"photo_path,omitempty"`
	RoleID          int        `json:"role_id"`
	Role            struct {
		RoleName  string `json:"role_name"`
		GuardName string `json:"guard_name"`
	}
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

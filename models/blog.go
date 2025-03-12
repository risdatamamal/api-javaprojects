package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Blog struct {
	GormModel
	BlogTitle  string `gorm:"not null;uniqueIndex" json:"blog_title" form:"blog_title" valid:"required~Blog title is required"`
	BlogDesc   string `gorm:"not null;uniqueIndex" json:"blog_desc" form:"blog_desc" valid:"required~Blog desc is required"`
	IsActive   bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath  string `json:"image_path" form:"image_path"`
	UserID     int    ` json:"user_id"`
	User       *User
	ViewsCount int `gorm:"null;default:0" json:"views_count" form:"views_count"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}

	return nil
}

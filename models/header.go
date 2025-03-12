package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Header struct {
	GormModel
	HeaderTitle string `gorm:"not null;uniqueIndex" json:"header_title" form:"header_title" valid:"required~Header title is required"`
	HeaderDesc  string `gorm:"not null;uniqueIndex" json:"header_desc" form:"header_desc" valid:"required~Header desc is required"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath   string `json:"image_path" form:"image_path"`
}

func (u *Header) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type About struct {
	GormModel
	AboutTitle string `gorm:"not null;uniqueIndex" json:"about_title" form:"about_title" valid:"required~About title is required"`
	AboutDesc  string `gorm:"not null;uniqueIndex" json:"about_desc" form:"about_desc" valid:"required~About desc is required"`
	IsActive   bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath  string `json:"image_path" form:"image_path"`
}

func (a *About) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	return nil
}

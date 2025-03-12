package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Industry struct {
	GormModel
	IndustryName string `gorm:"not null;uniqueIndex" json:"industry_name" form:"industry_name" valid:"required~Industry name is required"`
	IsActive     bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath    string `json:"image_path" form:"image_path"`
}

func (i *Industry) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(i)
	if err != nil {
		return err
	}

	return nil
}

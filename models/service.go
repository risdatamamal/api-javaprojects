package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Service struct {
	GormModel
	ServiceTitle string `gorm:"not null;uniqueIndex" json:"service_title" form:"service_title" valid:"required~Service title is required"`
	ServiceDesc  string `gorm:"not null;uniqueIndex" json:"service_desc" form:"service_desc" valid:"required~Service desc is required"`
	IsActive     bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath    string `json:"image_path" form:"image_path"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}

	return nil
}

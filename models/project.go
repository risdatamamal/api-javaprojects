package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Project struct {
	GormModel
	ProjectName string `gorm:"not null;uniqueIndex" json:"project_name" form:"project_name" valid:"required~Project name is required"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath   string `json:"image_path" form:"image_path"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

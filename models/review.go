package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Review struct {
	GormModel
	Content   string `gorm:"not null" json:"content" form:"content" valid:"required~Content is required"`
	Rating    int    `gorm:"not null" json:"rating" form:"rating" valid:"required~Rating is required"`
	ProjectID int    `json:"project_id"`
	Project   *Project
	UserID    int ` json:"user_id"`
	User      *User
}

func (p *Review) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	return nil
}

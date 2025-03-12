package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Client struct {
	GormModel
	ClientName string `gorm:"not null;uniqueIndex" json:"client_name" form:"client_name" valid:"required~Client name is required"`
	IsActive   bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath  string `json:"image_path" form:"image_path"`
}

func (c *Client) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}

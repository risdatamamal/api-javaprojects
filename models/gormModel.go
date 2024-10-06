package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (g *GormModel) BeforeCreate(tx *gorm.DB) (err error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(location)
	g.CreatedAt = &now
	g.UpdatedAt = &now
	return nil
}

func (g *GormModel) BeforeUpdate(tx *gorm.DB) (err error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(location)
	g.UpdatedAt = &now
	return nil
}

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
	currentDateTime, err := time.Parse(time.RFC3339, now.Format(time.RFC3339))

	if err != nil {
		return err
	}

	g.CreatedAt = &currentDateTime
	g.UpdatedAt = &currentDateTime
	return nil
}

func (g *GormModel) BeforeUpdate(tx *gorm.DB) (err error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(location)
	currentDateTime, err := time.Parse(time.RFC3339, now.Format(time.RFC3339))

	if err != nil {
		return err
	}

	g.UpdatedAt = &currentDateTime
	return nil
}

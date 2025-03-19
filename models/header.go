package models

import (
	"log"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Header struct {
	GormModel
	HeaderTitle string `gorm:"not null;uniqueIndex" json:"header_title" form:"header_title" valid:"required~Header title is required"`
	HeaderDesc  string `gorm:"not null;uniqueIndex" json:"header_desc" form:"header_desc" valid:"required~Header desc is required"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active" form:"is_active"`
	ImagePath   string `json:"image_path" form:"image_path"`
	LinkVideo   string `json:"link_video" form:"link_video"`
	FbLink      string `json:"fb_link" form:"fb_link"`
	YtLink      string `json:"yt_link" form:"yt_link"`
	IgLink      string `json:"ig_link" form:"ig_link"`
}

type GetAllHeadersResponse struct {
	GormModel
	HeaderTitle string `json:"header_title"`
	HeaderDesc  string `json:"header_desc"`
	IsActive    bool   `json:"is_active"`
	ImagePath   string `json:"image_path"`
	LinkVideo   string `json:"link_video"`
	FbLink      string `json:"fb_link"`
	YtLink      string `json:"yt_link"`
	IgLink      string `json:"ig_link"`
}

func (u *Header) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}

func SeedHeaders(db *gorm.DB) {
	headers := []Header{
		{
			HeaderTitle: "Providing superior IT solutions to ensure maximum contentment",
			HeaderDesc:  "At javaprojects we specialize in designing, building, shipping and scaling beautiful, usable products with blazing-fast efficiency",
			IsActive:    true,
			ImagePath:   "header1.jpg",
			LinkVideo:   "https://www.youtube.com/embed/4T7e4v4ZQ6A",
			FbLink:      "https://www.facebook.com",
			YtLink:      "https://www.youtube.com",
			IgLink:      "https://www.instagram.com",
		},
	}

	for _, header := range headers {
		var existingHeader Header

		if err := db.Where("header_title = ?", header.HeaderTitle).First(&existingHeader).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&header).Error; err != nil {
					log.Printf("Failed to seed header: %v", err)
				} else {
					log.Printf("Header %s has been created", header.HeaderTitle)
				}
			} else {
				log.Printf("Error checking header: %v", err)
			}
		} else {
			log.Printf("Header %s already exists", header.HeaderTitle)
		}
	}
}

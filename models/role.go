package models

import (
	"log"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Role struct {
	GormModel
	RoleName  string `gorm:"not null;uniqueIndex" json:"role_name" form:"role_name" valid:"required~Role name is required"`
	GuardName string `gorm:"not null" json:"guard_name" form:"guard_name" valid:"required~Guard name is required"`
}

type GetAllRolesResponse struct {
	GormModel
	RoleName  string `json:"role_name"`
	GuardName string `json:"guard_name"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		return err
	}

	return nil
}

func SeedRoles(db *gorm.DB) {
	roles := []Role{
		{RoleName: "Admin", GuardName: "web"},
		{RoleName: "User", GuardName: "web"},
	}

	for _, role := range roles {
		var existingRole Role

		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					log.Printf("Failed to seed role: %v", err)
				} else {
					log.Printf("Role %s has been created", role.RoleName)
				}
			} else {
				log.Printf("Error checking role: %v", err)
			}
		} else {
			log.Printf("Role %s already exists", role.RoleName)
		}
	}
}

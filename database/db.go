package database

import (
	"fmt"

	"github.com/risdatamamal/api-javaprojects/config"
	"github.com/risdatamamal/api-javaprojects/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() error {
	conf := config.LoadConfig()
	conn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%d sslmode=disable", conf.Host, conf.Username, conf.Password, conf.DBName, conf.Port)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("Successfully Connected to Database: ", conf.DBName)

	db.Debug().AutoMigrate(
		models.User{},
		models.Role{},
	// models.Review{},
	)

	models.SeedRoles(db)

	return nil
}

func GetDB() *gorm.DB {
	return db
}

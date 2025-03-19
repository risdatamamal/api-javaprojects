package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	SecretKey  string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid PORT value")
	}

	configs := Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		Host:       os.Getenv("HOST"),
		Port:       port,
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SecretKey:  os.Getenv("SECRET_KEY"),
	}

	fmt.Println("Username:", configs.Username)

	return configs
}

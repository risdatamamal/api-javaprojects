package config

import (
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
		Username:   os.Getenv("USERNAME"),
		Password:   os.Getenv("PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SecretKey:  os.Getenv("SECRET_KEY"),
	}

	return configs
}

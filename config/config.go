package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	DB DB
	HttpPort         string
}

type DB struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}


func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("no .env file, reading config from OS ENV variables")
	}
	return &Config{
		DB: DB{
			DBUser:           os.Getenv("DB_USER"),
			DBPass:           os.Getenv("DB_PASSWORD"),
			DBName:           os.Getenv("DB_NAME"),
			DBHost:           os.Getenv("DB_HOST"),
			DBPort:           os.Getenv("DB_PORT"),
		},
		HttpPort:         os.Getenv("HTTP_PORT"),
	}
}


package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port         string
	JWTSecret    string
	DBUser       string
	DBPass       string
	DBHost       string
	DBPort       string
	DBName       string
	CookieDomain string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		CookieDomain: os.Getenv("COOKIE_DOMAIN"),
	}
}

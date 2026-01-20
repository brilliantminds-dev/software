package config

import (
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	Host   string
	DBName string
	User   string
}

func GetAppConfig() *AppConfig {
	err := godotenv.Load("internal/config/.env")
	if err != nil {
		panic("error loading app config")
	}
	return &AppConfig{
		Host:   os.Getenv("HOST"),
		DBName: os.Getenv("DBNAME"),
		User:   os.Getenv("USER"),
	}

}

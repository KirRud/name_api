package config

import (
	"github.com/joho/godotenv"
	"name_api/internal/models"
	"os"
)

func InitConfig() (*models.Config, error) {
	var cfg models.Config
	err := godotenv.Load("./configs/dev.env")
	if err != nil {
		return nil, err
	}

	cfg.DB = models.DBConf{
		UserName: os.Getenv("DB_USERNAME"),
		Pass:     os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
	cfg.Server = models.ServerConf{
		Port: os.Getenv("SERVER_PORT"),
	}

	return &cfg, nil
}

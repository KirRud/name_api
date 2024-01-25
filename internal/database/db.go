package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"name_api/internal/models"
)

type DataBase struct {
	db *gorm.DB
}

type DataBaseInterface interface {
	personRepoInterface
	infoRepoInterface
	joinRepoInterface
}

func InitDB(conf models.DBConf) (DataBaseInterface, error) {
	conn := postgres.Open(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		conf.UserName, conf.Pass, conf.Database, conf.Host, conf.Port))

	db, err := gorm.Open(conn)
	if err != nil {
		log.Printf("Failed to initialize GORM: %v", err)
		return nil, err
	}

	if !db.Migrator().HasTable(&models.InfoDB{}) {
		db.Migrator().CreateTable(&models.InfoDB{})
	} else {
		db.AutoMigrate(&models.InfoDB{})
	}
	if !db.Migrator().HasTable(&models.PersonDB{}) {
		db.Migrator().CreateTable(&models.PersonDB{})
	} else {
		db.AutoMigrate(&models.PersonDB{})
	}

	sl := DataBase{db: db}
	return &sl, nil
}

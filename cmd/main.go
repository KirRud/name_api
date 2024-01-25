package main

import (
	log "github.com/sirupsen/logrus"
	"name_api/internal/config"
	controller2 "name_api/internal/controller"
	"name_api/internal/database"
	handler2 "name_api/internal/handler"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		return
	}

	db, err := database.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
		return
	}

	controller := controller2.NewController(db)
	handler := handler2.NewHandler(controller)
	router := handler2.InitRoutes(handler)

	router.Run(":" + cfg.Server.Port)
}

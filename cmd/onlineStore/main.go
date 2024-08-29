package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"onlineStore/internal/config"
	"onlineStore/internal/handlers"
	"onlineStore/internal/models"
	"onlineStore/internal/route"
	"os"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Http.Host)
	fmt.Println(cfg.Http.Port)

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Orders{}, &models.Payments{})

	h := handlers.New(db)

	r := route.Router(h)

	err = http.ListenAndServe(cfg.Http.Host+":"+cfg.Http.Port, r)
	if err != nil {
		return
	}

}

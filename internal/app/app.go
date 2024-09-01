package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"onlineStore/internal/config"
	"onlineStore/internal/handlers"
	"onlineStore/internal/models"
	"onlineStore/internal/repository"
	"onlineStore/internal/route"
)

func Run() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Http.Host)
	fmt.Println(cfg.Http.Port)

	//dsn := os.Getenv("POSTGRESQL_URL")
	dsn := "host=localhost user=postgres password=postgres dbname=example port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.Payment{})

	ur := repository.NewUserRepo(db)
	h := handlers.NewUserHandler(ur)

	r := route.Router(*h)

	err = http.ListenAndServe(cfg.Http.Host+":"+cfg.Http.Port, r)
	if err != nil {
		return
	}

}

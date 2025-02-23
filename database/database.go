package database

import (
	"fmt"
	"log"
	"os"

	"github.com/cli-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=Awaisniaz dbname=go_todo_backend port=5432 sslmode=disable"
	}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected Successfully!")
	DB.AutoMigrate(&models.Task{}, &models.User{})
}

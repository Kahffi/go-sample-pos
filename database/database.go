package database

import (
	"log"

	"github.com/Kahffi/go-sample-pos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Daffa123212321@tcp(localhost:8080)/go_sample_pos?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Product{})  // Migrasi tabel product
	database.AutoMigrate(&models.Receipt{})  // Migrasi tabel product
	database.AutoMigrate(&models.Payment{})  // Migrasi tabel product
	database.AutoMigrate(&models.Employee{}) // Migrasi tabel product
	database.AutoMigrate(&models.Order{})    // Migrasi tabel product
	database.AutoMigrate(&models.Customer{}) // Migrasi tabel product

	DB = database
}

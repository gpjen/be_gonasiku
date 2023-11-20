package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {

	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBHost, DBPort, DBDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("success to connect database")

	DB = db

	MigrateModels()
}

func MigrateModels() {
	DB.AutoMigrate()
}

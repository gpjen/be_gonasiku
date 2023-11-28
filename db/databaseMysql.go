package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDB() {

	dsn := os.Getenv("DB_MYSQL")
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

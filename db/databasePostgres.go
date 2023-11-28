package db

import (
	"fmt"
	"os"

	"gonasiku.com/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func InitPostgresDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_NAME"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("success to connect database")

	PostgresDB = db

	PostgresDB.AutoMigrate(&user.User{})
}

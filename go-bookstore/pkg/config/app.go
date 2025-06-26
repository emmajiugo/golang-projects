package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB
)

// ConnectDB initializes the database connection
func ConnectDB() {
	dsn := "user:secret123@tcp(127.0.0.1:3306)/gobookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	if DB == nil {
		panic("database connection is not initialized")
	}
	return DB
}
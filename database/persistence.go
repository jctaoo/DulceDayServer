package database

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB() *gorm.DB {
	dsn := config.SiteConfig.DataBaseConfig.GetDSN()
	fmt.Println("dsn: " + dsn)
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		fmt.Println("Some Error Occurred When Connect DB. ", dbErr)
	}
	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		err := db.AutoMigrate(&models.User{})
		if err != nil {
			fmt.Println("Some Error Occurred When Migrate Table")
		}
	} else {
		fmt.Println("Cannot set connection pool.")
	}
	return db
}
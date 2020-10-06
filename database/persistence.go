package database

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDB() *gorm.DB {
	dsn := config.SiteConfig.DataBaseConfig.GetDSN()
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		fmt.Println("Some Error Occurred When Connect DB. ", dbErr)
	}
	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		err := db.AutoMigrate(&models.User{}, &models.TokenAuth{})
		if err != nil {
			fmt.Println("Some Error Occurred When Migrate Table")
		}
	} else {
		fmt.Println("Cannot set connection pool.")
	}
	return db
}

func NewCache() *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.SiteConfig.CacheConfig.Host, config.SiteConfig.CacheConfig.Port)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB: config.SiteConfig.CacheConfig.DB,
		PoolSize: 10,
	})
	return client
}
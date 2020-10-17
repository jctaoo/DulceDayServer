package database

import (
	"DulceDayServer/config"
	"DulceDayServer/database/models"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewDB() *gorm.DB {
	dsn := config.SiteConfig.DataBaseConfig.GetDSN()
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		fmt.Println("Some Error Occurred When Connect DB. ", dbErr)
		os.Exit(-1)
	}
	if db != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		err := db.AutoMigrate(
			&models.User{},
			&models.TokenAuth{},
			&models.UserProfile{},
			&models.Moment{},
			&models.MomentStarUser{},
			&models.PurchaseItem{},
		)

		if err != nil {
			fmt.Println("Some Error Occurred When Migrate Table")
			os.Exit(-1)
		}
	} else {
		fmt.Println("Cannot set connection pool.")
		os.Exit(-1)
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

func NewAliOSS() *oss.Bucket {
	endpoint := config.SiteConfig.AliOssStaticStorageConfig.Endpoint
	keyId := config.SiteConfig.AliOssStaticStorageConfig.AccessKeyId
	secret := config.SiteConfig.AliOssStaticStorageConfig.AccessKeySecret
	bucketName := config.SiteConfig.AliOssStaticStorageConfig.BucketName

	client, err := oss.New(endpoint, keyId, secret)
	if err != nil {
		fmt.Println("Some Error Occurred When Create AliOSS Client. ", err)
		os.Exit(-1)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Some Error Occurred When Create AliOSS Client. ", err)
		os.Exit(-1)
	}
	return bucket
}

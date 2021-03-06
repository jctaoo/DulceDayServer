package database

import (
	"DulceDayServer/config"
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
		fmt.Println("Some Error Occurred When Connect DB server. ", dbErr)
		fmt.Printf("DSN is: %s\n", dsn)
		os.Exit(-1)
	} else {
		fmt.Printf("Connect DB successfully. DSN is: %s\n", dsn)
	}
	if db != nil {
		// 连接池
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	} else {
		fmt.Println("Cannot set connection pool for persistence DB.")
		os.Exit(-1)
	}
	return db
}

func NewCache() *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.SiteConfig.CacheConfig.Host, config.SiteConfig.CacheConfig.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       config.SiteConfig.CacheConfig.DB,
		PoolSize: 10,
	})
	fmt.Printf("Connected cache server. addr is: %s\n", addr)
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

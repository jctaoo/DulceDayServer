package database

import (
	"DulceDayServer/database/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"os"
)

// MigrateDataBase 用于数据迁移，当且仅当在项目启动时运行
func MigrateDataBase(db *gorm.DB, cdb *redis.Client) {
	var err error

	// todo 标记迁移状态，自动选择是否迁移

	// 常规迁移，每次运行前执行
	err = db.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.AuthUser{},
		&models.TokenAuth{},
		&models.Moment{},
		&models.MomentStarUser{},
	)
	if err != nil {
		fmt.Println("Some Error Occurred When Migrate Table")
		os.Exit(-1)
	}
}

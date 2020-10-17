package store

import (
	"DulceDayServer/database/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Repository interface {
	GetPurchases() *[]models.PurchaseItem
}

type RepositoryImpl struct {
	db *gorm.DB
	cdb *redis.Client
}

func NewRepositoryImpl(db *gorm.DB, cdb *redis.Client) *RepositoryImpl {
	return &RepositoryImpl{db: db, cdb: cdb}
}

func (r RepositoryImpl) GetPurchases() *[]models.PurchaseItem {
	var purchases []models.PurchaseItem
	r.db.Find(&purchases)
	return &purchases
}


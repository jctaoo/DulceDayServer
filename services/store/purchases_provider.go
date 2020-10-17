package store

import (
	"DulceDayServer/database/models"
)

// 内购商品提供者
// 用于获取商品信息，商品状态及其优惠信息等
type PurchasesProvider interface {
	GetPurchases() *[]models.PurchaseItem
}

type PurchasesProviderImpl struct {
	repo Repository
}

func NewPurchasesProviderImpl(repo Repository) *PurchasesProviderImpl {
	return &PurchasesProviderImpl{repo: repo}
}

func (p PurchasesProviderImpl) GetPurchases() *[]models.PurchaseItem {
	return p.repo.GetPurchases()
}

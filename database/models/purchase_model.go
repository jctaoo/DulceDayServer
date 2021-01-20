package models

import "gorm.io/gorm"

// 内购商品类型
type PurchaseType int

const (
	// 消耗型商品
	PurchaseTypeConsume PurchaseType = iota
	// 非自动续期订阅类商品
	PurchaseTypeSubscription
	// 自动续期订阅类商品
	PurchaseTypeAutomaticSubscription
)

// 内购商品模型
type PurchaseItem struct {
	gorm.Model `json:"-"`

	// 内部标示符，比如 AppStoreConnect 中的产品 ID
	InternalIdentifier string `json:"-"`

	// 内部名称
	InternalName string `json:"-"`

	// 商品类型，详见 PurchaseType
	Type PurchaseType `json:"type"`

	// 商品 ID
	Identifier string `json:"identifier"`
}

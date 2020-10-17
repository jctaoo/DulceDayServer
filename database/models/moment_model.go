package models

import (
	"DulceDayServer/helpers"
	"gorm.io/gorm"
)

// 图文动态的模型
type Moment struct {
	gorm.Model `json:"-"`

	// 使用 MomentID 字段来查找相应动态，
	// 形如 "M12138"
	// 使用类似 "/v1/moment/M12138" 的 Rest Api 来获取
	MomentID string `json:"moment_id"`

	// 发动态的用户的 Identifier
	UserIdentifier string `json:"-"`

	// 动态的文字内容
	Content string `json:"content"`
}

// 图文动态的点赞👍
type MomentStarUser struct {
	gorm.Model
	MomentID       string `json:"-"`
	UserIdentifier string `json:"-"`
}

func (m MomentStarUser) IsEmpty() bool {
	return m == MomentStarUser{}
}

func NewMoment(content string, userIdentifier string) *Moment {
	id := "M" + helpers.GenerateRandomKey()
	return &Moment{
		MomentID:       id,
		UserIdentifier: userIdentifier,
		Content:        content,
	}
}

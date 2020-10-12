package user_profile

import "DulceDayServer/database/models"

// 此处表示完整的用户模型
type FullUser struct {
	Username string `json:"username"`
	*models.UserProfile `json:"user_profile"`
}
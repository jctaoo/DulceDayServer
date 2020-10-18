package user_profile

import "DulceDayServer/database/models"

// FullUser 用于表示完整的用户信息
// Why:
// 单独的 models.UserProfile 无法表现完整的用户信息，比如 Username
// models.User 包含与用户信息不想关的信息
type FullUser struct {
	Username            string `json:"username"`
	*models.UserProfile `json:"profile"`
}

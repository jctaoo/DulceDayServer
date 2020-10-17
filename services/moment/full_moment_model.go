package moment

import "DulceDayServer/services/user_profile"

// 此处表示完整的动态模型
type FullMoment struct {
	MomentID string `json:"moment_id"`
	StarCount int64 `json:"star_count"`
	Content string `json:"content"`
	*user_profile.FullUser `json:"full_user"`
	Stared bool `json:"stared"`
}

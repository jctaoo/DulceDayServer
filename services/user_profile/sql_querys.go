package user_profile

import (
	"DulceDayServer/database/models"
	"gorm.io/gorm"
)

const kFullMomentSelectionQuery = `
	user_profiles.nickname, users.username,
 	users.user_identifier as user_identifier, user_profiles.avatar_file_key
`

const kFullMomentJoinQuery = `
	LEFT OUTER JOIN users on user_profiles.user_identifier = users.user_identifier
`

func buildBaseQueryForFullUser(db *gorm.DB) *gorm.DB {
	query := db.Model(&models.UserProfile{})
	query = query.Select(kFullMomentSelectionQuery)
	query = query.Joins(kFullMomentJoinQuery)
	return query
}

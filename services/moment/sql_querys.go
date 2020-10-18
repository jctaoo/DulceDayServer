package moment

import (
	"DulceDayServer/database/models"
	"gorm.io/gorm"
)

const kFullMomentSelectionQuery = `
	moments.id, moments.moment_id, moments.content, COUNT(moment_star_users.moment_id) AS star_count,
	user_profiles.nickname, users.username, user_profiles.avatar_file_key,
	users.user_identifier as user_identifier,
	IF(SUM(IF(moment_star_users.user_identifier = @login_user_identifier, 1, 0)) = true, true, false) as stared
`

const kFullMomentJoinQuery = `
	LEFT OUTER JOIN moment_star_users ON moments.moment_id = moment_star_users.moment_id
	LEFT OUTER JOIN user_profiles ON moments.user_identifier = user_profiles.user_identifier
	LEFT OUTER JOIN users ON moments.user_identifier = users.user_identifier
`

const kFullMomentGroupQuery = `
	moments.id, moment_star_users.moment_id, user_profiles.id, users.id
`

func buildBaseQueryForFullMoment(db *gorm.DB, loginUserIdentifier string) *gorm.DB {
	query := db.Model(&models.Moment{})
	query = query.Select(kFullMomentSelectionQuery, loginUserIdentifier)
	query = query.Joins(kFullMomentJoinQuery)
	query = query.Group(kFullMomentGroupQuery)
	return query
}

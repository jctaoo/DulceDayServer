package moment

import "gorm.io/gorm"

const kFullMomentSelectionQuery = `
	moments.id, moments.content, COUNT(moment_star_users.moment_id) AS star_count,
	user_profiles.uid, user_profiles.nickname, users.username, user_profiles.avatar_file_key,
	users.identifier as user_identifier,
	IF(SUM(IF(moment_star_users.user_identifier = ?, 1, 0)) = true, true, false) as stared
`

const kFullMomentJoinQuery = `
	LEFT OUTER JOIN moment_star_users ON moments.moment_id = moment_star_users.moment_id
	LEFT OUTER JOIN user_profiles ON moments.user_identifier = user_profiles.user_identifier
	LEFT OUTER JOIN users ON moments.user_identifier = users.identifier
`

const kFullMomentGroupQuery = `
	moments.id, moment_star_users.moment_id, user_profiles.id, users.id
`

func buildBaseQueryForFullMoment(db *gorm.DB, loginUserIdentifier string) *gorm.DB {
	query := db.Table("moments")
	query = query.Select(kFullMomentSelectionQuery, loginUserIdentifier)
	query = query.Joins(kFullMomentJoinQuery)
	query = query.Group(kFullMomentGroupQuery)
	return query
}
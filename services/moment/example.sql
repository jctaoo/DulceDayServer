use DulceDay;


# 查询所有动态生成 FullMoment
set @login_user_identifier = '123';
SELECT moments.id, moments.moment_id, moments.content, COUNT(moment_star_users.moment_id) AS star_count,
       user_profiles.nickname, users.username, user_profiles.avatar_file_key,
       users.user_identifier as user_identifier,
       IF(SUM(IF(moment_star_users.user_identifier = @login_user_identifier, 1, 0)) = true, true, false) as stared
FROM moments
LEFT OUTER JOIN moment_star_users ON moments.moment_id = moment_star_users.moment_id
LEFT OUTER JOIN user_profiles ON moments.user_identifier = user_profiles.user_identifier
LEFT OUTER JOIN users ON moments.user_identifier = users.user_identifier
GROUP BY moments.id, moment_star_users.moment_id, user_profiles.id, users.id;


# 为已登陆用户查询某条动态（通过MomentID），并且标示是否已经点赞，生成 FullMoment
set @login_user_identifier = '123';
set @moment_id = '3';
SELECT moments.id, moments.moment_id, moments.content, COUNT(moment_star_users.moment_id) AS star_count,
       user_profiles.nickname, users.username, user_profiles.avatar_file_key,
       users.user_identifier as user_identifier,
       IF(SUM(IF(moment_star_users.user_identifier = @login_user_identifier, 1, 0)) = true, true, false) as stared
FROM moments
LEFT OUTER JOIN moment_star_users ON moments.moment_id = moment_star_users.moment_id
LEFT OUTER JOIN user_profiles ON moments.user_identifier = user_profiles.user_identifier
LEFT OUTER JOIN users ON moments.user_identifier = users.user_identifier
WHERE moments.moment_id = @moment_id
GROUP BY moments.id, moment_star_users.moment_id, user_profiles.id, users.id;
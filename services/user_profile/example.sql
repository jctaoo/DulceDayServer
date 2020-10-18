USE DulceDay;

# 查询用户名为 jctaoo 的用户并生成 FullUser
set @user_name = 'jcatoo';
SELECT user_profiles.nickname, users.username,
 		users.user_identifier as user_identifier, user_profiles.avatar_file_key
FROM user_profiles
LEFT OUTER JOIN users on user_profiles.user_identifier = users.user_identifier
WHERE users.username = @user_name;
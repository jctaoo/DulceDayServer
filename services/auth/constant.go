package auth

// redis 中 ID BlackList 的 score 值常量
const kUserIdBlackListScore = 100

// redis 中 Token RevokeList 的 score 值常量
const kTokenRevokeListScore = 200

// redis 中 Token InActiveList 的 score 值常量
const kTokenInActiveListScore = 200

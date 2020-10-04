package user

// 密码错误
type ErrorPasswordWrong struct {

}

func (p ErrorPasswordWrong) Error() string {
	return "用户密码错误"
}

// ID 在黑名单中
type ErrorUserIdInBlackList struct {

}

func (I ErrorUserIdInBlackList) Error() string {
	return "该用户的ID已经被列入黑名单"
}

// Token 被 Revoke
type ErrorTokenRevoke struct {

}

func (e ErrorTokenRevoke) Error() string {
	return "Token 被撤回"
}

// Token is InActive
type ErrorTokenInActive struct {

}

func (e ErrorTokenInActive) Error() string {
	return "Token 未激活"
}

// Bad token
type ErrorTokenBad struct {}

func (e ErrorTokenBad) Error() string {
	return "Bad Token"
}





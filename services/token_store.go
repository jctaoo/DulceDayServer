package services

// 存储用户 Token 以支持灵活的基于 Token 的操作，如支持后期的可能出现的 OAuth 需求，以及 Token 的颁发与撤回功能
type TokenStore interface {
	AddNewToken(token string)
	FindToken(token string) bool
	RemoveToken(token string)
}

// 该实现采用 Redis 来实现 Token 的存储
type TokenStoreImpl struct {

}

func NewTokenStoreImpl() *TokenStoreImpl {
	return &TokenStoreImpl{}
}

func (t TokenStoreImpl) AddNewToken(token string) {

}

func (t TokenStoreImpl) FindToken(token string) bool {
	panic("implement me")
}

func (t TokenStoreImpl) RemoveToken(token string) {
	panic("implement me")
}

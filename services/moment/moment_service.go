package moment

type Service interface {
	// 创建新的 models.Moment
	// content 为文字内容，userIdentifier 为发动态的 models.User
	// 返回新创建的 models.Moment 的 MomentID
	CreateNewMoment(content string, userIdentifier string) (momentId string)

	// 切换给定 models.User 为给定 models.Moment 点赞状态
	// 根据当前是否点赞来判断相应动作
	ToggleStarMoment(momentId string, userIdentifier string) (isStarNow bool)

	// 通过 MomentID 查找 FullMoment
	// forUserIdentifier 是已登陆用户的 identifier 可空，用于查找给定用户与对应动态之间的状态
	// 比如给定用户是否为最终查找到的 models.Moment 点赞等
	GetMomentByMomentId(momentId string, forUserIdentifier string) *FullMoment

	// 在未登陆状态下获取推荐的 models.Moment
	GetRecommendMoments() *[]FullMoment

	// 在登陆状态下获取推荐的 models.Moment
	// 参数 userIdentifier 为已经登陆的用户的 identifier
	GetRecommendMomentsWithUserIdentifier(userIdentifier string) *[]FullMoment
}

type ServiceImpl struct {
	store Store
}

func NewServiceImpl(store Store) *ServiceImpl {
	return &ServiceImpl{store: store}
}

func (s ServiceImpl) CreateNewMoment(content string, userIdentifier string) (momentId string) {
	return s.store.createNewMoment(content, userIdentifier)
}

func (s ServiceImpl) ToggleStarMoment(momentId string, userIdentifier string) (isStarNow bool) {
	// 查询是否已经点过赞
	isStar := s.store.checkIsStar(momentId, userIdentifier)
	if isStar {
		// 点过赞，取消
		s.store.unStarMoment(momentId, userIdentifier)
		return false
	} else {
		// 未点赞，点赞
		s.store.starMoment(momentId, userIdentifier)
		return true
	}
}

func (s ServiceImpl) GetMomentByMomentId(momentId string, forUserIdentifier string) *FullMoment {
	return s.store.findMomentByMomentId(momentId, forUserIdentifier)
}

func (s ServiceImpl) GetRecommendMoments() *[]FullMoment {
	return s.store.findTheMomentsForUser("")
}

func (s ServiceImpl) GetRecommendMomentsWithUserIdentifier(userIdentifier string) *[]FullMoment {
	return s.store.findTheMomentsForUser(userIdentifier)
}

package moment

import (
	"DulceDayServer/database/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Store interface {
	// 创建新的 models.Moment
	// content 为文字内容，userIdentifier 为发动态的 models.User
	// 返回新创建的 models.Moment 的 MomentID
	createNewMoment(content string, userIdentifier string) (momentId string)

	// 给定 models.User 为某个 models.Moment 点赞, 插入相应 models.MomentStarUser 记录
	// momentId 标示要点赞的 models.Moment 的 MomentID
	// userIdentifier 标示要给该 models.Moment 点赞的 models.User
	starMoment(momentId string, userIdentifier string)

	// 给定 models.User 为某个 models.Moment 取消点赞, 删除(非软删除)相应 models.MomentStarUser 记录
	// momentId 标示要点赞的 models.Moment 的 MomentID
	// userIdentifier 标示要给该 models.Moment 点赞的 models.User
	unStarMoment(momentId string, userIdentifier string)

	// 检查给定 models.User 是否为给定 models.Moment 点赞
	// momentId 标示要点赞的 models.Moment 的 MomentID
	// userIdentifier 标示要给该 models.Moment 点赞的 models.User
	checkIsStar(momentId string, userIdentifier string) bool

	// 通过 MomentID 查找 FullMoment
	// forUserIdentifier 是已登陆用户的 identifier 可空，用于查找给定用户与对应动态之间的状态
	// 比如给定用户是否为最终查找到的 models.Moment 点赞等
	findMomentByMomentId(momentId string, forUserIdentifier string) *FullMoment

	// 通过各种条件为某个用户寻找 models.Moment
	// 用例: 全站最火的十条图文动态
	// 参数 userIdentifier 是可选的，如果未提供该参数，则视为为游客用户寻找对应图文动态
	findTheMomentsForUser(userIdentifier string) *[]FullMoment
}

type StoreImpl struct {
	db *gorm.DB
	cdb *redis.Client
}

func NewStoreImpl(db *gorm.DB, cdb *redis.Client) *StoreImpl {
	return &StoreImpl{db: db, cdb: cdb}
}

func (s StoreImpl) createNewMoment(content string, userIdentifier string) (momentId string) {
	moment := models.NewMoment(content, userIdentifier)
	s.db.Create(moment)
	return moment.MomentID
}

func (s StoreImpl) starMoment(momentId string, userIdentifier string) {
	if momentId == "" || userIdentifier == "" {
		return
	}
	star := &models.MomentStarUser{
		MomentID:       momentId,
		UserIdentifier: userIdentifier,
	}
	s.db.Create(star)
}

func (s StoreImpl) unStarMoment(momentId string, userIdentifier string) {
	if momentId == "" || userIdentifier == "" {
		return
	}
	s.db.Unscoped().Delete(&models.MomentStarUser{MomentID: momentId, UserIdentifier: userIdentifier})
}

func (s StoreImpl) checkIsStar(momentId string, userIdentifier string) bool {
	star := &models.MomentStarUser{}
	s.db.Where(models.MomentStarUser{MomentID: momentId, UserIdentifier: userIdentifier}).Take(star)
	return !star.IsEmpty()
}

func (s StoreImpl) findMomentByMomentId(momentId string, forUserIdentifier string) *FullMoment {
	moment := &FullMoment{}
	query := buildBaseQueryForFullMoment(s.db, forUserIdentifier)
	query = query.Where(&models.Moment{
		MomentID: momentId,
	})
	query.Take(moment)
	return moment
}

func (s StoreImpl) findTheMomentsForUser(userIdentifier string) *[]FullMoment {
	var resFullMoments []FullMoment
	// todo 此时只是找全部的 moment，未实现定制推送
	buildBaseQueryForFullMoment(s.db, userIdentifier).Find(&resFullMoments)
	return &resFullMoments
}

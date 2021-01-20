package moment

import (
	"DulceDayServer/api/common"
	serviceAuth "DulceDayServer/services/auth"
	"DulceDayServer/services/moment"
	"DulceDayServer/services/static_storage"
	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	common.BaseEndpoints
	// 为路人获取推荐的动态
	requestRecommendMoments(context *gin.Context)
	// 为已经登陆的用户精准推送推荐的动态
	requestRecommendMomentsWithAuth(context *gin.Context)
	// 获取某个动态的详细信息
	getMoment(context *gin.Context)
	// 创建动态
	createMoment(context *gin.Context)
	// 删除动态
	deleteMoment(context *gin.Context)
	// 更新动态信息(可见性等)
	updateMoment(context *gin.Context)
	// 更改点赞👍
	toggleMomentStar(context *gin.Context)
}

type EndpointsImpl struct {
	service       moment.Service
	userService   serviceAuth.Service
	staticStorage static_storage.Service
}

func NewEndpointsImpl(service moment.Service, userService serviceAuth.Service, staticStorage static_storage.Service) *EndpointsImpl {
	return &EndpointsImpl{service: service, userService: userService, staticStorage: staticStorage}
}

func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	group := router.Group("/moment")

	group.GET("/recommend/hot", e.requestRecommendMoments)
	group.GET(
		"/recommend",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.requestRecommendMomentsWithAuth,
	)

	group.GET(
		"/get/:MomentID",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyAccess),
		e.getMoment,
	)

	group.POST(
		"/create",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.createMoment,
	)

	group.PATCH(
		"/toggle_star/:MomentID",
		common.MiddleWareAuth(e.userService, common.MiddleWareAuthPolicyReject),
		e.toggleMomentStar,
	)

	return group
}

func (e EndpointsImpl) deleteMoment(context *gin.Context) {
	panic("implement me")
}

func (e EndpointsImpl) updateMoment(context *gin.Context) {
	panic("implement me")
}

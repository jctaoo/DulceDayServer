package graphql

import (
	"DulceDayServer/api/common"
	"DulceDayServer/graphql/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"path"
)

type Endpoints interface {
	common.BaseEndpoints
}

type EndpointsImpl struct {
	Endpoints
}

func NewEndpointsImpl() *EndpointsImpl {
	return &EndpointsImpl{}
}

// TODO add to config.toml
func (e EndpointsImpl) MapHandlersToRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	group := router.Group("/graphql")

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &Resolver{},
			},
		),
	)

	queryPath := path.Join(group.BasePath(), "query")
	ginSrv := gin.WrapH(srv)
	ginPlayground := gin.WrapH(playground.Handler("outer-graphql", queryPath))

	group.GET("/", ginPlayground)
	group.POST("/query", ginSrv)

	return group
}

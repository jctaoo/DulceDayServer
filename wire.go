package main
//
//import (
//	"DulceDayServer/api"
//	"DulceDayServer/api/user"
//	"DulceDayServer/database"
//	"DulceDayServer/services"
//	"github.com/google/wire"
//)
//
//var userEndpointSet = wire.NewSet(
//	services.NewAuthService,
//	services.NewGenericService,
//	database.NewDB,
//	api.NewHttpStatusPackage,
//	wire.Bind(new(api.HttpPackage), new(*api.HttpPackageImpl)),
//	wire.Bind(new(services.AuthService), new(*services.AuthServiceImpl)),
//	wire.Bind(new(services.GenericService), new(*services.GenericServiceImpl)),
//)
//
//func UserEndpoints() user.Endpoints {
//	panic(
//		wire.Build(
//			user.NewEndpointsImpl,
//			userEndpointSet,
//			wire.Bind(new(user.Endpoints), new(*user.EndpointsImpl)),
//		),
//	)
//}
//
//func HttpStatusPackage() api.HttpPackage {
//	panic(
//		wire.Build(
//			api.NewHttpStatusPackage,
//			wire.Bind(new(api.HttpPackage), new(*api.HttpPackageImpl)),
//		),
//	)
//}
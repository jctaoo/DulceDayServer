// +build wireinject

package main

import (
	"DulceDayServer/api/base"
	"DulceDayServer/api/user"
	"DulceDayServer/database"
	"github.com/google/wire"
	serviceUser "DulceDayServer/services/user"
)

var userEndpointsSet = wire.NewSet(
	database.NewCache,
	database.NewDB,

	serviceUser.NewServiceImpl,
	wire.Bind(new(serviceUser.Service), new(*serviceUser.ServiceImpl)),

	serviceUser.NewEncryptionAdaptorImpl,
	wire.Bind(new(serviceUser.EncryptionAdaptor), new(*serviceUser.EncryptionAdaptorImpl)),

	serviceUser.NewTokenGranterImpl,
	wire.Bind(new(serviceUser.TokenGranter), new(*serviceUser.TokenGranterImpl)),

	serviceUser.NewStoreImpl,
	wire.Bind(new(serviceUser.Store), new(*serviceUser.StoreImpl)),

	serviceUser.NewTokenStoreImpl,
	wire.Bind(new(serviceUser.TokenStore), new(*serviceUser.TokenStoreImpl)),

	serviceUser.NewTokenAdaptorImpl,
	wire.Bind(new(serviceUser.TokenAdaptor), new(*serviceUser.TokenAdaptorImpl)),

	base.NewHttpStatusPackage,
	wire.Bind(new(base.HttpPackage), new(*base.HttpPackageImpl)),
)

func UserEndpoints() user.Endpoints {
	panic(
		wire.Build(
			user.NewEndpointsImpl,
			userEndpointsSet,
			wire.Bind(new(user.Endpoints), new(*user.EndpointsImpl)),
		),
	)
}

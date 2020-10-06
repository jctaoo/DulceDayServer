// +build wireinject

package main

import (
	"DulceDayServer/api/user"
	"DulceDayServer/api/user_profile"
	"DulceDayServer/database"
	serviceUser "DulceDayServer/services/user"
	serviceUserProfile "DulceDayServer/services/user_profile"
	"github.com/google/wire"
)

var userServiceEndpointSet = wire.NewSet(
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
)

var userEndpointsSet = wire.NewSet(
	database.NewCache,
	database.NewDB,

	userServiceEndpointSet,
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

var userProfileEndpointsSet = wire.NewSet(
	database.NewCache,
	database.NewDB,

	serviceUserProfile.NewUserProfileServiceImpl,
	wire.Bind(new(serviceUserProfile.Service), new(*serviceUserProfile.ServiceImpl)),

	serviceUserProfile.NewUserProfileStoreImpl,
	wire.Bind(new(serviceUserProfile.UserProfileStore), new(*serviceUserProfile.UserProfileStoreImpl)),

	userServiceEndpointSet,
)

func UserProfileEndpoints() user_profile.Endpoints {
	panic(
		wire.Build(
			user_profile.NewEndpointsImpl,
			userProfileEndpointsSet,
			wire.Bind(new(user_profile.Endpoints), new(*user_profile.EndpointsImpl)),
		),
	)
}


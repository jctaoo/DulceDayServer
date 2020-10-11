// +build wireinject

package main

import (
	"DulceDayServer/api/moment"
	"DulceDayServer/api/user"
	"DulceDayServer/api/user_profile"
	"DulceDayServer/database"
	"DulceDayServer/services/static_storage"
	apiStaticStorage "DulceDayServer/api/static_storage"
	serviceUser "DulceDayServer/services/user"
	serviceUserProfile "DulceDayServer/services/user_profile"
	serviceMoment "DulceDayServer/services/moment"
	"github.com/google/wire"
)

var universalSet = wire.NewSet(database.NewCache, database.NewDB, database.NewAliOSS)

var aliossStaticStorageServiceSet = wire.NewSet(
	static_storage.NewAliOSSStaticStorageService,
	wire.Bind(new(static_storage.Service), new(*static_storage.AliOSSStaticStorageService)),
)

var userServiceSet = wire.NewSet(
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

var userProfileServiceSet = wire.NewSet(
	serviceUserProfile.NewUserProfileServiceImpl,
	wire.Bind(new(serviceUserProfile.Service), new(*serviceUserProfile.ServiceImpl)),

	serviceUserProfile.NewUserProfileStoreImpl,
	wire.Bind(new(serviceUserProfile.UserProfileStore), new(*serviceUserProfile.UserProfileStoreImpl)),
)

var userEndpointsSet = wire.NewSet(
	universalSet,
	userServiceSet,
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
	universalSet,
	userServiceSet,
	userProfileServiceSet,
	aliossStaticStorageServiceSet,
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

var staticStorageEndpointsSet = wire.NewSet(
	universalSet,
	aliossStaticStorageServiceSet,
)

func StaticStorageEndpoints() apiStaticStorage.Endpoints {
	panic(
		wire.Build(
			apiStaticStorage.NewEndpointsImpl,
			staticStorageEndpointsSet,
			wire.Bind(new(apiStaticStorage.Endpoints), new(*apiStaticStorage.EndpointsImpl)),
		),
	)
}

var momentServiceSet = wire.NewSet(
	serviceMoment.NewServiceImpl,
	wire.Bind(new(serviceMoment.Service), new(*serviceMoment.ServiceImpl)),

	serviceMoment.NewStoreImpl,
	wire.Bind(new(serviceMoment.Store), new(*serviceMoment.StoreImpl)),
)

var momentEndpointsSet = wire.NewSet(
	universalSet,
	userServiceSet,
	momentServiceSet,
	aliossStaticStorageServiceSet,
)

func MomentEndpoints() moment.Endpoints {
	panic(
		wire.Build(
			moment.NewEndpointsImpl,
			momentEndpointsSet,
			wire.Bind(new(moment.Endpoints), new(*moment.EndpointsImpl)),
		),
	)
}
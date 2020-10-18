// +build wireinject

package main

import (
	"DulceDayServer/api/moment"
	apiStaticStorage "DulceDayServer/api/static_storage"
	"DulceDayServer/api/store"
	"DulceDayServer/api/user"
	"DulceDayServer/api/user_profile"
	"DulceDayServer/database"
	serviceMoment "DulceDayServer/services/moment"
	"DulceDayServer/services/static_storage"
	serviceUser "DulceDayServer/services/user"
	serviceAuth "DulceDayServer/services/auth"
	serviceUserProfile "DulceDayServer/services/user_profile"
	serviceMoment "DulceDayServer/services/moment"
	serviceStore "DulceDayServer/services/store"
	"github.com/google/wire"
)

var universalSet = wire.NewSet(database.NewCache, database.NewDB, database.NewAliOSS)

var aliossStaticStorageServiceSet = wire.NewSet(
	static_storage.NewAliOSSStaticStorageService,
	wire.Bind(new(static_storage.Service), new(*static_storage.AliOSSStaticStorageService)),
)

var authUserServiceSet = wire.NewSet(
	serviceAuth.NewServiceImpl,
	wire.Bind(new(serviceAuth.Service), new(*serviceAuth.ServiceImpl)),

	serviceAuth.NewEncryptionAdaptorImpl,
	wire.Bind(new(serviceAuth.EncryptionAdaptor), new(*serviceAuth.EncryptionAdaptorImpl)),

	serviceAuth.NewTokenGranterImpl,
	wire.Bind(new(serviceAuth.TokenGranter), new(*serviceAuth.TokenGranterImpl)),

	serviceAuth.NewStoreImpl,
	wire.Bind(new(serviceAuth.Store), new(*serviceAuth.StoreImpl)),

	serviceAuth.NewTokenStoreImpl,
	wire.Bind(new(serviceAuth.TokenStore), new(*serviceAuth.TokenStoreImpl)),

	serviceAuth.NewTokenAdaptorImpl,
	wire.Bind(new(serviceAuth.TokenAdaptor), new(*serviceAuth.TokenAdaptorImpl)),
)

var userServiceSet = wire.NewSet(
	serviceUser.NewStoreImpl,
	wire.Bind(new(serviceUser.Store), new(*serviceUser.StoreImpl)),

	serviceUser.NewServiceImpl,
	wire.Bind(new(serviceUser.Service), new(*serviceUser.ServiceImpl)),
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
	userProfileServiceSet,
	authUserServiceSet,
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
	authUserServiceSet,
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
	authUserServiceSet,
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


var storeServiceSet = wire.NewSet(
	serviceStore.NewPurchasesProviderImpl,
	wire.Bind(new(serviceStore.PurchasesProvider), new(*serviceStore.PurchasesProviderImpl)),

	serviceStore.NewRepositoryImpl,
	wire.Bind(new(serviceStore.Repository), new(*serviceStore.RepositoryImpl)),
)

var storeEndpointsSet = wire.NewSet(
	universalSet,
	storeServiceSet,
)

func StoreEndpoints() store.Endpoints {
	panic(
		wire.Build(
			store.NewEndpointsImpl,
			storeEndpointsSet,
			wire.Bind(new(store.Endpoints), new(*store.EndpointsImpl)),
		),
	)
}
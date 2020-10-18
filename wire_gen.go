// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"DulceDayServer/api/moment"
	static_storage2 "DulceDayServer/api/static_storage"
	"DulceDayServer/api/store"
	"DulceDayServer/api/user"
	user_profile2 "DulceDayServer/api/user_profile"
	"DulceDayServer/database"
	"DulceDayServer/services/auth"
	moment2 "DulceDayServer/services/moment"
	"DulceDayServer/services/static_storage"
	store2 "DulceDayServer/services/store"
	user2 "DulceDayServer/services/user"
	"DulceDayServer/services/user_profile"
	"github.com/google/wire"
)

import (
	_ "DulceDayServer/docs"
)

// Injectors from wire.go:

func UserEndpoints() user.Endpoints {
	db := database.NewDB()
	client := database.NewCache()
	storeImpl := user2.NewStoreImpl(db, client)
	encryptionAdaptorImpl := auth.NewEncryptionAdaptorImpl()
	tokenStoreImpl := auth.NewTokenStoreImpl(db, client)
	tokenAdaptorImpl := auth.NewTokenAdaptorImpl()
	tokenGranterImpl := auth.NewTokenGranterImpl(tokenStoreImpl, tokenAdaptorImpl)
	authStoreImpl := auth.NewStoreImpl(db, client)
	serviceImpl := auth.NewServiceImpl(encryptionAdaptorImpl, tokenGranterImpl, authStoreImpl)
	userProfileStoreImpl := user_profile.NewUserProfileStoreImpl(db, client)
	user_profileServiceImpl := user_profile.NewUserProfileServiceImpl(userProfileStoreImpl)
	userServiceImpl := user2.NewServiceImpl(storeImpl, serviceImpl, user_profileServiceImpl)
	endpointsImpl := user.NewEndpointsImpl(userServiceImpl, serviceImpl)
	return endpointsImpl
}

func UserProfileEndpoints() user_profile2.Endpoints {
	db := database.NewDB()
	client := database.NewCache()
	userProfileStoreImpl := user_profile.NewUserProfileStoreImpl(db, client)
	serviceImpl := user_profile.NewUserProfileServiceImpl(userProfileStoreImpl)
	encryptionAdaptorImpl := auth.NewEncryptionAdaptorImpl()
	tokenStoreImpl := auth.NewTokenStoreImpl(db, client)
	tokenAdaptorImpl := auth.NewTokenAdaptorImpl()
	tokenGranterImpl := auth.NewTokenGranterImpl(tokenStoreImpl, tokenAdaptorImpl)
	storeImpl := auth.NewStoreImpl(db, client)
	authServiceImpl := auth.NewServiceImpl(encryptionAdaptorImpl, tokenGranterImpl, storeImpl)
	bucket := database.NewAliOSS()
	aliOSSStaticStorageService := static_storage.NewAliOSSStaticStorageService(bucket)
	endpointsImpl := user_profile2.NewEndpointsImpl(serviceImpl, authServiceImpl, aliOSSStaticStorageService)
	return endpointsImpl
}

func StaticStorageEndpoints() static_storage2.Endpoints {
	bucket := database.NewAliOSS()
	aliOSSStaticStorageService := static_storage.NewAliOSSStaticStorageService(bucket)
	endpointsImpl := static_storage2.NewEndpointsImpl(aliOSSStaticStorageService)
	return endpointsImpl
}

func MomentEndpoints() moment.Endpoints {
	db := database.NewDB()
	client := database.NewCache()
	storeImpl := moment2.NewStoreImpl(db, client)
	serviceImpl := moment2.NewServiceImpl(storeImpl)
	encryptionAdaptorImpl := auth.NewEncryptionAdaptorImpl()
	tokenStoreImpl := auth.NewTokenStoreImpl(db, client)
	tokenAdaptorImpl := auth.NewTokenAdaptorImpl()
	tokenGranterImpl := auth.NewTokenGranterImpl(tokenStoreImpl, tokenAdaptorImpl)
	authStoreImpl := auth.NewStoreImpl(db, client)
	authServiceImpl := auth.NewServiceImpl(encryptionAdaptorImpl, tokenGranterImpl, authStoreImpl)
	bucket := database.NewAliOSS()
	aliOSSStaticStorageService := static_storage.NewAliOSSStaticStorageService(bucket)
	endpointsImpl := moment.NewEndpointsImpl(serviceImpl, authServiceImpl, aliOSSStaticStorageService)
	return endpointsImpl
}

func StoreEndpoints() store.Endpoints {
	db := database.NewDB()
	client := database.NewCache()
	repositoryImpl := store2.NewRepositoryImpl(db, client)
	purchasesProviderImpl := store2.NewPurchasesProviderImpl(repositoryImpl)
	endpointsImpl := store.NewEndpointsImpl(purchasesProviderImpl)
	return endpointsImpl
}

// wire.go:

var universalSet = wire.NewSet(database.NewCache, database.NewDB, database.NewAliOSS)

var aliossStaticStorageServiceSet = wire.NewSet(static_storage.NewAliOSSStaticStorageService, wire.Bind(new(static_storage.Service), new(*static_storage.AliOSSStaticStorageService)))

var authUserServiceSet = wire.NewSet(auth.NewServiceImpl, wire.Bind(new(auth.Service), new(*auth.ServiceImpl)), auth.NewEncryptionAdaptorImpl, wire.Bind(new(auth.EncryptionAdaptor), new(*auth.EncryptionAdaptorImpl)), auth.NewTokenGranterImpl, wire.Bind(new(auth.TokenGranter), new(*auth.TokenGranterImpl)), auth.NewStoreImpl, wire.Bind(new(auth.Store), new(*auth.StoreImpl)), auth.NewTokenStoreImpl, wire.Bind(new(auth.TokenStore), new(*auth.TokenStoreImpl)), auth.NewTokenAdaptorImpl, wire.Bind(new(auth.TokenAdaptor), new(*auth.TokenAdaptorImpl)))

var userServiceSet = wire.NewSet(user2.NewStoreImpl, wire.Bind(new(user2.Store), new(*user2.StoreImpl)), user2.NewServiceImpl, wire.Bind(new(user2.Service), new(*user2.ServiceImpl)))

var userProfileServiceSet = wire.NewSet(user_profile.NewUserProfileServiceImpl, wire.Bind(new(user_profile.Service), new(*user_profile.ServiceImpl)), user_profile.NewUserProfileStoreImpl, wire.Bind(new(user_profile.UserProfileStore), new(*user_profile.UserProfileStoreImpl)))

var userEndpointsSet = wire.NewSet(
	universalSet,
	userServiceSet,
	userProfileServiceSet,
	authUserServiceSet,
)

var userProfileEndpointsSet = wire.NewSet(
	universalSet,
	authUserServiceSet,
	userProfileServiceSet,
	aliossStaticStorageServiceSet,
)

var staticStorageEndpointsSet = wire.NewSet(
	universalSet,
	aliossStaticStorageServiceSet,
)

var momentServiceSet = wire.NewSet(moment2.NewServiceImpl, wire.Bind(new(moment2.Service), new(*moment2.ServiceImpl)), moment2.NewStoreImpl, wire.Bind(new(moment2.Store), new(*moment2.StoreImpl)))

var momentEndpointsSet = wire.NewSet(
	universalSet,
	authUserServiceSet,
	momentServiceSet,
	aliossStaticStorageServiceSet,
)

var storeServiceSet = wire.NewSet(store2.NewPurchasesProviderImpl, wire.Bind(new(store2.PurchasesProvider), new(*store2.PurchasesProviderImpl)), store2.NewRepositoryImpl, wire.Bind(new(store2.Repository), new(*store2.RepositoryImpl)))

var storeEndpointsSet = wire.NewSet(
	universalSet,
	storeServiceSet,
)

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"DulceDayServer/api/base"
	"DulceDayServer/api/user"
	"DulceDayServer/database"
	"DulceDayServer/services"
	"github.com/google/wire"
)

// Injectors from wire.go:

func UserEndpoints() user.Endpoints {
	db := database.NewDB()
	authServiceImpl := services.NewAuthService(db)
	genericServiceImpl := services.NewGenericService(db)
	httpPackageImpl := base.NewHttpStatusPackage()
	endpointsImpl := user.NewEndpointsImpl(authServiceImpl, genericServiceImpl, httpPackageImpl)
	return endpointsImpl
}

func HttpStatusPackage() base.HttpPackage {
	httpPackageImpl := base.NewHttpStatusPackage()
	return httpPackageImpl
}

// wire.go:

var userEndpointSet = wire.NewSet(services.NewAuthService, services.NewGenericService, database.NewDB, base.NewHttpStatusPackage, wire.Bind(new(base.HttpPackage), new(*base.HttpPackageImpl)), wire.Bind(new(services.AuthService), new(*services.AuthServiceImpl)), wire.Bind(new(services.GenericService), new(*services.GenericServiceImpl)))
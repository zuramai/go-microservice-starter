package usecasefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/logger"
)

var UseCaseFactoryBuilderMap = map[string]UseCaseFbInterface{
	config.REGISTRATION: &RegistrationFactory{},
	config.LIST_USER:    &ListUserFactory{},
	config.UPDATE_USER:  &UpdateUserFactory{},
	config.DELETE_USER:  &DeleteUserFactory{},
}

type UseCaseInterface interface{}

type UseCaseFbInterface interface {
	Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error)
}

//GetDataStoreFb is accessors for factoryBuilderMap
func GetUseCaseFb(key string) UseCaseFbInterface {
	v, ok := UseCaseFactoryBuilderMap[key]
	if !ok {
		logger.SugarLog.Error("Use case not found in UseCaseFactoryBuilderMap{}")
	}
	return v
}

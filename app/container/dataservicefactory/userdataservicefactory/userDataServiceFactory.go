package userdataservicefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/applicationservice/dataservice"
)

var userDataServiceFbMap = map[string]userDataServiceFbInterface{
	config.MONGODB: &mongodbUserDataServiceFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type userDataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (dataservice.UserDataInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetUserDataServiceFb(key string) userDataServiceFbInterface {
	return userDataServiceFbMap[key]
}

package datastorefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
)

var dataStoreFactoryMap = map[string]dataStoreFbInterface{
	config.MONGODB:    &mongoFactory{},
	config.CACHE_GRPC: &cacheGrpcFactory{},
}

type DataStoreInterface interface{}

type dataStoreFbInterface interface {
	Build(container.Container, *config.DataStoreConfig) (DataStoreInterface, error)
}

func GetDataStoreFb(code string) dataStoreFbInterface {
	return dataStoreFactoryMap[code]
}

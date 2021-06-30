package servicecontainer

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container/usecasefactory"
	"github.com/muka-id/service-user/app/logger"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	AppConfig  *config.AppConfig
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	logger.SugarLog.Debug("Building Usecase with key: ", code)
	usecase, err := usecasefactory.GetUseCaseFb(code).Build(sc, sc.AppConfig, code)
	if err != nil {
		logger.SugarLog.Error("Error building use case", err)
		return nil, errors.Wrap(err, "Error building use case ")
	}
	return usecase, nil
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}

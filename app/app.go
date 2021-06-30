package app

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/container/servicecontainer"
	"github.com/muka-id/service-user/app/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func InitApp(filename ...string) (container.Container, error) {
	config, err := config.BuildConfig(filename...)
	if err != nil {
		return nil, errors.Wrap(err, "BuildConfig")
	}
	err = initLogger(config.ZapConfig)
	if err != nil {
		return nil, err
	}

	return initContainer(config)

}

func initContainer(config *config.AppConfig) (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{FactoryMap: factoryMap, AppConfig: config}
	return &c, nil
}

func initLogger(zc zap.Config) error {
	err := logger.SetLogger(zc)
	if err != nil {
		return err
	}
	logger.SugarLog.Debug("logger construction succeed")
	return nil
}

package dataservicefactory

import (
	"github.com/muka-id/service-user/app/config"
	"github.com/muka-id/service-user/app/container"
	"github.com/muka-id/service-user/app/container/dataservicefactory/userdataservicefactory"
	"github.com/muka-id/service-user/app/logger"
	"github.com/pkg/errors"
)

type userDataServiceFactory struct{}

func (userdatasf *userDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.SugarLog.Debugf("UserDataServiceFactory")
	key := dataConfig.DataStoreConfig.Code

	userdatasi, err := userdataservicefactory.GetUserDataServiceFb(key).Build(c, dataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return userdatasi, nil
}
